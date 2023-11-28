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
	Reports []Report
	Tests   map[string][]string
}

func WriteHTML(reports []Report, filename string) error {
	slog.Info("writing html report")

	testmap := make(map[string]map[string]bool)
	for _, report := range reports {
		for category, tests := range report.Results {
			if _, ok := tests[category]; !ok {
				testmap[category] = map[string]bool{}
			}

			for test := range tests {
				testmap[category][test] = true
			}
		}
	}

	templateInput := htmlTemplateInput{
		Reports: reports,
		Tests:   make(map[string][]string),
	}

	for category, tests := range testmap {
		for test := range tests {
			templateInput.Tests[category] = append(templateInput.Tests[category], test)
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
