# web5 conformance report

SDK: [{{ .TestServerID.Name }}]({{ .TestServerID.Url }}) ({{ .TestServerID.Language }})

| Test | Pass | Details |
| ---- | ---- | ------- |{{ range $test, $result := .Results }}
| `{{ $test }}` | {{ if $result }}:x: | ```{{ $result }}```{{ else }}:heavy_check_mark: |{{ end }} |{{ end }}
