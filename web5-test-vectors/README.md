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

Test vector files should adhere to [`vectors.schema.json`]('./vectors.schema.json'). This repo contains a [`test-vector-validation`]('../scripts/test-vector-validation') script that validates all vectors files in this directory. It can be run manually by following the directions in the script directory's README.

> [!NOTE]
> Vector Validation runs automatically anytime a change is made in this directory or to the script itself.

Each test vector is a structured collection of test vectors, where each vector is designed to validate specific features or behaviors. The schema ensures that the file includes a general description of the collection and an array of test vectors. Each test vector within the array provides a detailed account of what is being tested, including the input data, expected output data, and whether an error is anticipated.

Below is a table that outlines the expected fields in the JSON file:

| Field                   | Type    | Description                                                                                          | Required |
| ----------------------- | ------- | ---------------------------------------------------------------------------------------------------- | :------: |
| `description`           | string  | A general description of the test vectors collection.                                                |   Yes    |
| `vectors`               | array   | An array of test vectors for testing different features.                                             |   Yes    |
| `vectors[].description` | string  | A description of what this test vector is validating.                                                |   Yes    |
| `vectors[].input`       | any     | The input for the test vector, which can be of any type.                                             |   Yes    |
| `vectors[].output`      | any     | The expected output for the test vector, which can be of any type.                                   |    No    |
| `vectors[].errors`      | boolean | Indicates whether the test vector is expected to produce an error. Defaults to false if not present. |    No    |

Each `vectors[]` element represents a single test scenario, complete with its own description, inputs, and expected results, which are used to validate the corresponding feature or functionality.
