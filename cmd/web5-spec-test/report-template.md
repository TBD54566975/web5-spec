# web5 conformance report

SDK: [{{ .TestServerID.Name }}]({{ .TestServerID.Url }}) ({{ .TestServerID.Language }})

| Test | Pass | Details |
| ---- | ---- | ------- |{{ range $test, $result := .Results }}
| `{{ $test }}` | {{ if $result }}:heavy_check_mark: | {{ else }}:x: | -```{{ $result }}{{ end }} |{{ end }}
