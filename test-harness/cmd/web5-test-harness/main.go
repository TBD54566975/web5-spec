package main

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var (
	root = cobra.Command{
		Use: "web5-spec",
	}

	dockerfiles = []string{
		".web5-spec/Dockerfile",
		"web5-spec.Dockerfile",
	}
)

func main() {
	root.Execute()
}

func docker(dir string, args ...string) *exec.Cmd {
	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	slog.Info("executing docker", "args", args)

	return cmd
}
