# web5 conformance report <!-- markdownlint-disable -->

SDK: [{{ .TestServerID.Name }}]({{ .TestServerID.Url }}) ({{ .TestServerID.Language }})

{{ range $groupName, $results := .Results }}

## {{ $groupName }}

| Feature | Result |
| ------- | ------ |{{ range $test, $result := $results }}
| {{ $test }} | {{ $result.GetEmoji }}{{ if $result.Errors }} <ul>{{ range $result.Errors }}<li><pre>{{ . | sanatizeHTML }}</pre></li>{{ end }}</ul>{{ end }} |{{ end }}

{{ end }}
