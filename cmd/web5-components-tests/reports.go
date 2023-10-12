package main

import (
	"strings"

	"github.com/TBD54566975/web5-spec/openapi"
)

type Report struct {
	TestServerID openapi.TestServerID
	Results      map[string]error
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

func (r Report) Pass() bool {
	for _, err := range r.Results {
		if err != nil {
			return false
		}
	}

	return true
}
