# SDK Feature Report Generator

This tool reads junit XML reports from different SDKs and generates an HTML report showing which SDKs support which features.

## Current Status

`./cmd/build-html` will iterate over the repos listed in `sdk.go` and download the most recent junit artifact. It will read
all junit results from it and produce a report to `_site/index.html`
