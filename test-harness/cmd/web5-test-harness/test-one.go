package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/TBD54566975/sdk-development/openapi"
	"github.com/TBD54566975/sdk-development/reports"
	"github.com/TBD54566975/sdk-development/tests"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var (
	testOneCmd = &cobra.Command{
		Use: "one [dir]",
		Run: func(cmd *cobra.Command, args []string) {
			dir, _ := os.Getwd()
			if len(args) > 0 {
				dir = args[0]
			}

			report, err := testOne(dir)
			if err != nil {
				panic(err)
			}

			fmt.Println()
			if txt, err := report.Text(); err != nil {
				slog.Error("error generating text report", "error", err)
			} else {
				fmt.Println(txt)
			}
			fmt.Println()

			stepSummaryFile := os.Getenv("GITHUB_STEP_SUMMARY")
			if stepSummaryFile != "" {
				if err := reports.WriteMarkdown(report, stepSummaryFile); err != nil {
					slog.Error("error writing github step summary", "file", stepSummaryFile, "error", err)
				}
			}

			buildkiteToken := os.Getenv("BUILDKITE_ANALYTICS_TOKEN")
			if buildkiteToken != "" {
				if err := buildkiteUpload(report, buildkiteToken); err != nil {
					slog.Error("error sending results to buildkite", "error", err)
				}
			}

			junitFile := os.Getenv("JUNIT_REPORT")
			if junitFile != "" {
				if err := reports.WriteJunitToFile(report, junitFile); err != nil {
					slog.Error("error writing junit report", "file", junitFile, "error", err)
				}
			}

			if !report.IsPassing() {
				os.Exit(1)
			}

			os.Exit(0)
		},
	}
)

func testOne(dir string) (reports.Report, error) {
	logger := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	slog.SetDefault(slog.New(logger))

	var dockerfile string
	for _, d := range dockerfiles {
		candidate := filepath.Join(dir, d)
		if _, err := os.Stat(candidate); !os.IsNotExist(err) {
			dockerfile = d
		}
	}

	if dockerfile == "" {
		return reports.Report{}, fmt.Errorf("no dockerfile found (paths=%v)", dockerfiles)
	}

	cmd := docker(dir, "build", "-t", "web5-spec:latest", "-f", dockerfile, ".")
	if err := cmd.Run(); err != nil {
		return reports.Report{}, fmt.Errorf("error building server: %v", err)
	}

	cmd = docker(dir, "run", "-p", "8080:8080", "--name", "web5-spec", "--rm", "web5-spec:latest")
	if err := cmd.Start(); err != nil {
		return reports.Report{}, fmt.Errorf("error running server: %v", err)
	}

	defer func() {
		cmd := docker(dir, "stop", "web5-spec")
		if err := cmd.Run(); err != nil {
			slog.Error("error stopping server container", "error", err)
		}
	}()

	ctx := context.Background()

	client, err := openapi.NewClientWithResponses("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	var serverID openapi.TestServerID
	for {
		serverIDResponse, err := client.IdentifySelfWithResponse(ctx)
		if err != nil {
			slog.Debug("waiting for server to come up", "err", err)
			time.Sleep(time.Second)
			continue
		}

		if serverIDResponse.JSON200 == nil {
			slog.Debug("server ID check failed, retrying in 1 second", "status", serverIDResponse.Status(), "body", string(serverIDResponse.Body))
			time.Sleep(time.Second)
			continue
		}

		serverID = *serverIDResponse.JSON200
		break
	}

	defer func() {
		_, err := client.ServerShutdown(context.Background())
		if err != nil {
			slog.Error("error shutting down server", "error", err)
		}
	}()

	slog.Debug("server running", "sdk", serverID.Name, "url", serverID.Url)

	return reports.Report{
		TestServerID: serverID,
		Results:      tests.RunTests("http://localhost:8080"),
	}, nil
}

func init() {
	root.AddCommand(testOneCmd)
}

func buildkiteUpload(report reports.Report, buildkiteToken string) error {
	slog.Info("uploading junit report to buildkit")

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	formFields := map[string]string{
		"format":              "junit",
		"run_env[CI]":         "github_actions",
		"run_env[key]":        fmt.Sprintf("%s-%s-%s", os.Getenv("GITHUB_ACTION"), os.Getenv("GITHUB_RUN_NUMBER"), os.Getenv("GITHUB_RUN_ATTEMPT")),
		"run_env[number]":     os.Getenv("GITHUB_RUN_NUMBER"),
		"run_env[branch]":     os.Getenv("GITHUB_REF"),
		"run_env[commit_sha]": os.Getenv("GITHUB_SHA"),
		"run_env[url]":        fmt.Sprintf("https://github.com/%s/actions/runs/%s", os.Getenv("GITHUB_REPOSITORY"), os.Getenv("GITHUB_RUN_ID")),
	}
	for k, v := range formFields {
		if err := addFormValue(writer, k, v); err != nil {
			return err
		}
	}

	part, err := writer.CreateFormFile("data", "web5-test-harness.xml")
	if err != nil {
		slog.Error("error creating form file")
		return err
	}

	if err := reports.WriteJunit(report, part); err != nil {
		slog.Error("error generating junit report")
		return err
	}

	if err := writer.Close(); err != nil {
		slog.Error("error closing multi-part form writer")
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "https://analytics-api.buildkite.com/v1/uploads", &body)
	if err != nil {
		slog.Error("error constructing request to buildkite")
		return err
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", buildkiteToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("error uploading results to buildkite")
		return err
	}
	defer resp.Body.Close()

	responseBody := bytes.Buffer{}
	if _, err := io.Copy(&responseBody, resp.Body); err != nil {
		slog.Error("error reading response from buildkite")
		return err
	}

	if resp.StatusCode > 299 {
		slog.Error("unexpected response status from buildkite", "status", resp.Status, "response", body.String())
		return fmt.Errorf("unexpected %s", resp.Status)
	}

	slog.Info("successfully uploaded results to buildkite", "response", body.String())

	return nil
}

func addFormValue(writer *multipart.Writer, key, value string) error {
	field, err := writer.CreateFormField(key)
	if err != nil {
		slog.Error("error creating form field", "key", key, "value", value)
		return err
	}

	_, err = io.Copy(field, strings.NewReader(value))
	if err != nil {
		slog.Error("error writing form field value", "key", key, "value", value)
		return err
	}

	return nil
}
