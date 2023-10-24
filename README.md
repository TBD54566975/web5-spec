# web5-spec

This repo houses the specification and test suite for Web5 SDKs.

| Description                                |                                                                                |
| ------------------------------------------ | ------------------------------------------------------------------------------ |
| [Specification](./spec/README.md)          | Web5 SDK specification                                                         |
| [CODEOWNERS](./CODEOWNERS)                 | Outlines the project lead(s)                                                   |
| [CODE_OF_CONDUCT.md](./CODE_OF_CONDUCT.md) | Expected behavior for project contributors, promoting a welcoming environment  |
| [CONTRIBUTING.md](./CONTRIBUTING.md)       | Developer guide to build, test, run, access CI, chat, discuss, file issues     |
| [GOVERNANCE.md](./GOVERNANCE.md)           | Project governance                                                             |
| [LICENSE](./LICENSE)                       | Apache License, Version 2.0                                                    |


To generate the golang open api models:
```bash
make
```

To generate the javascript open api models:
```bash
npx openapi-typescript openapi.yaml -o sdks/web5-js/openapi.d.ts
```

To generate the kotlin open api models:
```bash
# Note: After running the above command, you will need to manually copy the generated model files into the correct directory in the sdk/web5-kt models directory and delete the generated folder.
openapi-generator generate -i openapi.yaml -g kotlin-server -o ./kotlin-server-generated-server --skip-validate-spec
```

To run the tests for javascript:

```bash
go run ./cmd/web5-spec-test sdks/web5-js
```

To run the tests for kotlin:

```bash
go run ./cmd/web5-spec-test sdks/web5-kt
```