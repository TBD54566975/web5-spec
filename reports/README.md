# SDK Feature Report Generator

This tool reads junit XML reports from different SDKs and generates an HTML report showing which SDKs support which features.

`./cmd/build-html` will iterate over the repos listed in `sdk.go` and download the most recent junit artifact. It will read all junit results from it and produce a report to `_site/index.html`.  For local testing, generate a
[GitHub Personal Access Token](https://github.com/settings/tokens?type=beta) and put it in an environment variable named
`GITHUB_TOKEN`. It doesn't need any special permissions ("Public Repositories (read-only)"). Note that a GitHub app (explained below) can also be used.

`./cmd/sync-vectors` will check the `main` branch of all SDKs listed in `sdks.go` and ensure their vectors match the ones in this repo.
For local testing, a [GitHub App](https://github.com/settings/apps) must be created. Put it's credentials in the following environment variables:

* `CICD_ROBOT_GITHUB_APP_ID` - shown on the edit page of the app, where you are sent right after app creation.
* `CICD_ROBOT_GITHUB_APP_PRIVATE_KEY` - this should be the contents of the private key, not the path to the file.
* `CICD_ROBOT_GITHUB_APP_NAME` - this is used as a display name and should match the name in the URL of the edit page for the app.
* `CICD_ROBOT_GITHUB_APP_INSTALLATION_ID` - click "Install App" on the sidebar while editing the app in GitHub to install it on your own account.

## Tooling

This project uses [hermit](https://cashapp.github.io/hermit/usage/get-started/), an open source toolchain manager, which pins and automatically downloads and installs tooling for a repo, including compiler toolchains, utilities, etc.

To install hermit, run:

```bash
https://github.com/cashapp/hermit/releases/download/stable/install.sh | /bin/bash
```

If using goland or intellij, also install the hermit plugin via [these instructions](https://cashapp.github.io/hermit/usage/ide/).
