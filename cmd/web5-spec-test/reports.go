package main

import (
	"embed"
	"html/template"
	"os"
	"strings"

	"github.com/TBD54566975/web5-spec/openapi"
)

//go:embed report-template.md
var reportTemplate embed.FS

type Report struct {
	TestServerID openapi.TestServerID
	Results      map[string]error
}

func (r Report) IsPassing() bool {
	for _, err := range r.Results {
		if err != nil {
			return false
		}
	}

	return true
}

func (r Report) Text() string {
	var b strings.Builder

	b.WriteString("web5 spec conformance report for ")
	b.WriteString(r.TestServerID.Name)
	b.WriteString(" (")
	b.WriteString(r.TestServerID.Url)
	b.WriteRune(')')
	b.WriteRune('\n')
	b.WriteRune('\n')
	for name, err := range r.Results {
		b.WriteString(name)
		b.WriteString(": ")
		if err != nil {
			b.WriteString("fail (")
			b.WriteString(err.Error())
			b.WriteRune(')')
		} else {
			b.WriteString("pass")
		}
	}

	return b.String()
}

var mdTemplate = template.Must(template.New("").ParseFS(reportTemplate, "report-template.md"))

func (r Report) WriteMarkdown(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = mdTemplate.ExecuteTemplate(f, "report-template.md", r)
	if err != nil {
		return err
	}

	return nil
}
