package reports

import "os"

func WriteMarkdown(report Report, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = textTemplates.ExecuteTemplate(f, "report-template.md", report)
	if err != nil {
		return err
	}

	return nil
}
