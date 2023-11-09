# Web5-kt Local Testing and Model Generation

## Running Tests Locally

To run the web5-kt tests on your local machine, execute the following command:

```bash
go run ./cmd/web5-spec-test sdks/web5-kt
```

## Generating Model Files
To generate the model files, use the openapi-generator with the following command:

```bash
openapi-generator generate -i openapi.yaml -g kotlin-server -o ./kotlin-server-generated-server --skip-validate-spec
```

Note: After running the above command, you will need to manually copy the generated model files into the correct directory in the sdk/web5-kt models directory.