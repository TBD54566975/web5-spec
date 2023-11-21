# Web5 Test Vectors

## Description

This directory contains test vectors for all features we intend to support accross several languages. Each feature has its own directory which contains a single vectors file per sub-feature e.g.

```text
web5-test-vectors
├── README.md
├── did-jwk <--- feature
│   └── resolve.json <--- sub-feature
├── index.html
└── vectors.schema.json
```

## Test Vector Files

Test vector files should adhere to [`vectors.schema.json`]('./vectors.schema.json'). This repo contains a [`test-vector-validation`]('../scripts/test-vector-validation') script that validates all vectors files in this directory. It can be run manually by following the instructions [here](../scripts/test-vector-validation/README.md).

> [!NOTE]
> Test Vector Validation runs automatically anytime a change is made in this directory or to the script itself.

Each test vector file is a structured collection of test vector objects, where each vector asserts a specific outcome for a given input. Below is a table that outlines the expected fields in a test vector file:

| Field                   | Type    | Description                                                                                          | Required |
| ----------------------- | ------- | ---------------------------------------------------------------------------------------------------- | :------: |
| `description`           | string  | A general description of the test vectors collection.                                                |   Yes    |
| `vectors`               | array   | An array of test vector objects.                                                                     |   Yes    |
| `vectors[].description` | string  | A description of what this test vector is testing.                                                   |   Yes    |
| `vectors[].input`       | any     | The input for the test vector, which can be of any type.                                             |   Yes    |
| `vectors[].output`      | any     | The expected output for the test vector, which can be of any type.                                   |    No    |
| `vectors[].errors`      | boolean | Indicates whether the test vector is expected to produce an error. Defaults to false if not present. |    No    |

### Rationale for Test Vector Structure

The structure of a `vector` object is designed to fulfill two conditions:

* the function works and returns something that should match `output`
* the function throws an error (in whatever way the consuming language represents errors)
  * _optionally_, the error's _output_ should match `output`

`errors: true` is an instruction to anticipate an error in the implementation language. For example:

* In languages like Kotlin or Javascript, the presence of `errors: true` would imply that `assertThrows` be used.
* In Go, the expectation would be for the err variable to be non-nil.
* In Rust, the error handling would pivot on matching `Result.Err` rather than `Result.Ok`.

Should `errors` be set to `true`, the `output` field may optionally be used to include expected error messages.
