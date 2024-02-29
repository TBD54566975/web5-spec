# `VC JWT` Test Vectors

This directory contains test vectors for the vc jwt functionality
[JWT](https://datatracker.ietf.org/doc/html/rfc7519).

## `verify`

Verify test vectors are detailed in a [JSON file](./verify.json).

### Input

The `input` for the sign operation is an object with the following properties:

| Property | Description                                                          |
| -------- | -------------------------------------------------------------------- |
| `jwt`    | A JSON Web Token JWT object representing the jwt to verify.          |

### Output

The `output` for the verify operation should not exist, it throws an error or it does not. To test validity of error can check the error message