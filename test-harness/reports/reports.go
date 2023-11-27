package reports

import (
	"embed"
	"errors"
	htmltemplate "html/template"
	texttemplate "text/template"
	"time"

	"github.com/TBD54566975/sdk-development/openapi"
)

var (
	ErrNotSupported = errors.New("test not supported by this SDK")

	//go:embed *
	templatesFS embed.FS

	htmlTemplates = htmltemplate.New("")
	textTemplates = texttemplate.New("")
	funcmap       = map[string]any{
		"sanatizeHTML":    sanatizeHTML,
		"durationToJunit": durationToJunit,
	}
)

func init() {
	htmlTemplates.Funcs(funcmap)
	if _, err := htmlTemplates.ParseFS(templatesFS, "report-template.html"); err != nil {
		panic(err)
	}

	textTemplates.Funcs(funcmap)
	if _, err := textTemplates.ParseFS(templatesFS, "report-template.*"); err != nil {
		panic(err)
	}
}

type Report struct {
	TestServerID openapi.TestServerID
	Results      map[string]map[string]Result
}

type Result struct {
	Errors []error
	Time   time.Duration
}

func (r Report) IsPassing() bool {
	for _, results := range r.Results {
		for _, result := range results {
			if result.IsSkipped() {
				continue
			}

			if len(result.Errors) > 0 {
				return false
			}
		}

	}

	return true
}

func (r Result) IsSkipped() bool {
	return len(r.Errors) == 1 && r.Errors[0] == ErrNotSupported
}

func (r Result) GetEmoji() string {
	if len(r.Errors) == 0 {
		return "✅"
	}

	if len(r.Errors) == 1 && r.Errors[0] == ErrNotSupported {
		return "🚧"
	}

	return "❌"
}
