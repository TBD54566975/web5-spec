# web5 conformance report

SDK: [{{ .TestServerID.Name }}]({{ .TestServerID.Url }}) ({{ .TestServerID.Language }})

{{ range $groupName, $results := .Results }}

## {{ $groupName }}

| Feature | Result |
| ------- | ------ |{{ range $test, $result := $results }}
| {{ $test }} | {{ if $result }}:x: <ul>{{ range $result }}<li><pre>{{ . }}</pre></li>{{ end }}</ul>{{ else }}:heavy_check_mark:{{ end }} |{{ end }}

{{ end }}
