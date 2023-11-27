package reports

import (
	"fmt"
	"io"
	"os"
	"time"
)

func durationToJunit(t time.Duration) string {
	return fmt.Sprintf("%f", float64(t)/float64(time.Second))
}

type junitTemplateInput struct {
	TimeTotal        time.Duration
	TimePerTestSuite map[string]time.Duration
	Report           Report
}

func WriteJunitToFile(report Report, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := WriteJunit(report, f); err != nil {
		return err
	}

	return nil
}

func WriteJunit(report Report, writer io.Writer) error {
	templateInput := junitTemplateInput{
		Report:           report,
		TimePerTestSuite: map[string]time.Duration{},
	}

	for testsuite, results := range report.Results {
		var testsuiteTime time.Duration
		for _, result := range results {
			testsuiteTime = testsuiteTime + result.Time
			templateInput.TimeTotal = templateInput.TimeTotal + result.Time
		}

		templateInput.TimePerTestSuite[testsuite] = testsuiteTime
	}

	if err := textTemplates.ExecuteTemplate(writer, "report-template.junit.xml", templateInput); err != nil {
		return err
	}

	return nil
}
