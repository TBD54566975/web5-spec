# SDK Feature Report Generator

This tool reads junit XML reports from different SDKs and generates an HTML report showing which SDKs support which features.

## Current Status

`./cmd/build-html` will iterate over the repos listed in `sdk.go` and download the most recent junit artifact. It will read
all junit results from it and produce a report to `_site/index.html`

## Tooling

This project uses [hermit](https://cashapp.github.io/hermit/usage/get-started/), an open source toolchain manager, which pins and automatically downloads and installs tooling for a repo, including compiler toolchains, utilities, etc.

To install hermit, run:

```bash
https://github.com/cashapp/hermit/releases/download/stable/install.sh | /bin/bash
```

If using goland or intellij, also install the hermit plugin via [these instructions](https://cashapp.github.io/hermit/usage/ide/).