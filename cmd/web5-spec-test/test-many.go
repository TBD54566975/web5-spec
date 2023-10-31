package main

import (
	"fmt"
	"os"

	"github.com/TBD54566975/web5-spec/reports"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var (
	testManyCmd = &cobra.Command{
		Use:  "many dir [dir...]",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			allReports := []reports.Report{}
			for _, dir := range args {
				report, err := testOne(dir)
				if err != nil {
					slog.Error("error testing server", "dir", dir, "err", err)
					continue
				}

				allReports = append(allReports, report)
			}

			for _, report := range allReports {
				fmt.Println()
				if txt, err := report.Text(); err != nil {
					slog.Error("error generating text report", "error", err)
					continue
				} else {
					fmt.Println(txt)
				}
				fmt.Println()
			}

			if err := os.MkdirAll("_site", 0755); err != nil && err != os.ErrExist {
				slog.Error("error creating _site/ for HTML report")
				panic(err)
			}
			if err := reports.WriteHTML(allReports, "_site/index.html"); err != nil {
				slog.Error("error rendering HTML template", "err", err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	root.AddCommand(testManyCmd)
}
