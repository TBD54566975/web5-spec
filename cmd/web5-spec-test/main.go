package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/TBD54566975/web5-spec/openapi"
	"github.com/TBD54566975/web5-spec/tests"
	"golang.org/x/exp/slog"
)

var (
	nostart = flag.Bool("no-start", false, "when set, the server is not built and is expected to be already running")
	nostop  = flag.Bool("no-stop", false, "when set, the server is not asked to shut down")
	server  = flag.String("server", "http://localhost:8080", "url of the server to connect to")
)

func main() {
	flag.Parse()

	dir, _ := os.Getwd()
	if len(flag.Args()) > 0 {
		dir = flag.Arg(0)
	}

	logger := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	slog.SetDefault(slog.New(logger))

	if !*nostart {
		cmd := run(dir, "docker", "build", "-t", "web5-component:latest", "-f", ".web5-component/test.Dockerfile", ".")
		if err := cmd.Run(); err != nil {
			slog.Error("error building server", "error", err)
			os.Exit(1)
		}

		cmd = run(dir, "docker", "run", "-p", "8080:8080", "--rm", "web5-component:latest")
		if err := cmd.Start(); err != nil {
			slog.Error("error running server", "error", err)
			os.Exit(1)
		}

		if !*nostop {
			defer func() {
				slog.Debug("shutting down server")
				if err := cmd.Process.Signal(os.Kill); err != nil {
					slog.Error("error shutting down server", "error", err)
				}
				if err := cmd.Wait(); err != nil {
					slog.Error("error shutting down server", "error", err)
				}
			}()
		}
	}

	ctx := context.Background()

	client, err := openapi.NewClientWithResponses(*server)
	if err != nil {
		panic(err)
	}

	var serverID openapi.TestServerID
	for {
		serverIDResponse, err := client.IdentifySelfWithResponse(ctx)
		if err != nil {
			slog.Debug("waiting for server to be ready", "err", err)
			time.Sleep(time.Second)
			continue
		}

		if serverIDResponse.JSON200 == nil {
			slog.Debug("server ID check failed", "status", serverIDResponse.Status(), "body", string(serverIDResponse.Body))
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

	slog.Debug("running tests")

	report := Report{
		TestServerID: serverID,
		Results:      tests.RunTests(*server),
	}

	fmt.Println()
	fmt.Println(report.Text())
	fmt.Println()

	stepSummaryFile := os.Getenv("GITHUB_STEP_SUMMARY")
	if stepSummaryFile != "" {
		if err := report.WriteMarkdown(stepSummaryFile); err != nil {
			slog.Error("error writing github step summary", "file", stepSummaryFile, "error", err)
		}
	}

	if !report.IsPassing() {
		os.Exit(1)
	}
}

func run(dir, command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	slog.Info("executing", "cmd", command, "args", args)

	return cmd
}
