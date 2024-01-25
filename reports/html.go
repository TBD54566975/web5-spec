package reports

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slog"
)

func sanatizeHTML(dirty error) string {
	clean := strings.ReplaceAll(dirty.Error(), "<", "&lt;")
	clean = strings.ReplaceAll(clean, ">", "&gt;")
	clean = strings.ReplaceAll(clean, "\n", "\\\\n")

	return clean
}

type htmlTemplateInput struct {
	Reports    []Report
	Web5Tests  map[string][]string
	TbDEXTests map[string][]string
}

func WriteHTML(reports []Report, filename string) error {
	slog.Info("writing html report", "reports", len(reports))

	testmap := make(map[string]map[string]bool)
	tbdexTestMap := make(map[string]map[string]bool)
	for _, report := range reports {
		for category, tests := range report.Results {
			if _, ok := tests[category]; !ok {
				if report.SDK.Type == "web5" {
					testmap[category] = map[string]bool{}
				} else {
					tbdexTestMap[category] = map[string]bool{}
				}
			}

			for test := range tests {
				if report.SDK.Type == "web5" {
					testmap[category][test] = true
				} else {
					tbdexTestMap[category][test] = true
				}
			}
		}
	}

	templateInput := htmlTemplateInput{
		Reports:    reports,
		Web5Tests:  make(map[string][]string),
		TbDEXTests: make(map[string][]string),
	}

	for category, tests := range testmap {
		for test := range tests {
			templateInput.Web5Tests[category] = append(templateInput.Web5Tests[category], test)
		}
	}

	for category, tests := range tbdexTestMap {
		for test := range tests {
			templateInput.TbDEXTests[category] = append(templateInput.TbDEXTests[category], test)
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error opening %s: %v", filename, err)
	}
	defer f.Close()

	if err := htmlTemplates.ExecuteTemplate(f, "report-template.html", templateInput); err != nil {
		return err
	}

	return nil
}
