package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/TBD54566975/web5-spec/openapi"
	"github.com/TBD54566975/web5-spec/reports"
	"github.com/TBD54566975/web5-spec/tests"
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
