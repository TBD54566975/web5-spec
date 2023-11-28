package reports

import "bytes"

func (r Report) Text() (string, error) {
	var buffer bytes.Buffer

	if err := textTemplates.ExecuteTemplate(&buffer, "report-template.txt", r); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
