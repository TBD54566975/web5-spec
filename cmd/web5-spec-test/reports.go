package main

import (
	"bytes"
	"embed"
	"os"
	"text/template"

	"github.com/TBD54566975/web5-spec/openapi"
)

//go:embed report-template.*
var reportTemplateFS embed.FS

var templates = template.Must(template.New("").ParseFS(reportTemplateFS, "report-template.*"))

type Report struct {
	TestServerID openapi.TestServerID
	Results      map[string]map[string][]error
}

func (r Report) IsPassing() bool {
	for _, errs := range r.Results {
		if len(errs) > 0 {
			return false
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

func (r Report) WriteMarkdown(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = templates.ExecuteTemplate(f, "report-template.md", r)
	if err != nil {
		return err
	}

	return nil
}
