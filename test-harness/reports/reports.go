package reports

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"os"
	"strings"

	"golang.org/x/exp/slog"

	"github.com/TBD54566975/web5-spec/openapi"
	"github.com/TBD54566975/web5-spec/tests"
)

//go:embed *
var templatesFS embed.FS

var templates = template.New("")

func init() {
	templates.Funcs(template.FuncMap{
		"sanatizeHTML": sanatizeHTML,
		"getEmoji":     getEmoji,
	})
	_, err := templates.ParseFS(templatesFS, "report-template.*")
	if err != nil {
		panic(err)
	}
}

type Report struct {
	TestServerID openapi.TestServerID
	Results      map[string]map[string][]error
}

func (r Report) IsPassing() bool {
	for _, results := range r.Results {
		for _, errs := range results {
			if len(errs) == 1 && errs[0] == tests.ErrNotSupported {
				continue
			}

			if len(errs) > 0 {
				return false
			}
		}

	}

	return true
}

func (r Report) Text() (string, error) {
	var buffer bytes.Buffer

	if err := templates.ExecuteTemplate(&buffer, "report-template.txt", r); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func WriteMarkdown(report Report, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = templates.ExecuteTemplate(f, "report-template.md", report)
	if err != nil {
		return err
	}

	return nil
}

func sanatizeHTML(dirty error) string {
	clean := strings.ReplaceAll(dirty.Error(), "<", "&lt;")
	clean = strings.ReplaceAll(clean, ">", "&gt;")
	clean = strings.ReplaceAll(clean, "\n", "\\\\n")

	return clean
}

func getEmoji(errs []error) string {
	if len(errs) == 0 {
		return "✅"
	}

	if len(errs) == 1 && errs[0] == tests.ErrNotSupported {
		return "🚧"
	}

	return "❌"
}

type htmlTemplateInput struct {
	Reports []Report
	Tests   map[string][]string
}

func WriteHTML(reports []Report, filename string) error {
	slog.Info("writing html report")

	testmap := map[string]map[string]bool{}
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

	templateInput := htmlTemplateInput{Reports: reports, Tests: map[string][]string{}}

	for category, tests := range testmap {
		templateInput.Tests[category] = []string{}
		for test := range tests {
			templateInput.Tests[category] = append(templateInput.Tests[category], test)
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error opening %s: %v", filename, err)
	}
	defer f.Close()

	if err := templates.ExecuteTemplate(f, "report-template.html", templateInput); err != nil {
		return err
	}

	return nil
}
