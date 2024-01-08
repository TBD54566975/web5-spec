package main

import (
	"os"

	"github.com/TBD54566975/sdk-development/reports"
	"golang.org/x/exp/slog"
)

func main() {
	defer reports.CleanupGitAuth()
	if err := reports.ConfigureGitAuth(); err != nil {
		panic(err)
	}

	errs := make(map[string]error)
	for _, sdk := range reports.SDKs {
		if err := reports.SyncSDK(sdk); err != nil {
			errs[sdk.Name] = err
		}
	}

	if err := reports.CleanupGitAuth(); err != nil {
		panic(err)
	}

	if len(errs) > 0 {
		for sdk, err := range errs {
			slog.Error("error", "sdk", sdk, "error", err)
		}
		os.Exit(1)
	}
}
