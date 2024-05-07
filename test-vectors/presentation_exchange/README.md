# Presentation Exchange Test Vectors

The reference implementation for Presentation Exchange can be found [here](https://github.com/TBD54566975/web5-js/blob/main/packages/credentials/src/presentation-exchange.ts#L80).

## `create_presentation_from_credentials.json`
Vectors for creating a Presentation Submission from a given Presentation Definition and set of VC JWTs.

| Property                        | Description                                                                                                                                                                                              |
|---------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `input.presentationDefinition`  | the input [presentationDefinition](https://identity.foundation/presentation-exchange/#presentation-definition) showing the requirements used for this getting an example WA license                      |
| `input.credentialJwt`           | the input [verifiable credential secured as a JWT](https://www.w3.org/TR/vc-data-model/#json-web-token) that corresponds to the presentationDefinition to fulfill it and do a full presentation exchange |
| `output.presentationSubmission` | the expected [presentationSubmission](https://identity.foundation/presentation-exchange/#presentation-submission) when the `inputs` are processed by the `createPresentationFromCredentials` method.     |

## `select_credentials.json`
Vectors for selecting a subset of VC JWTs from a given set that satisfy a given Presentation Definition

| Property                       | Description                                                                                                                                                                                                                                                          |
|--------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `input.presentationDefinition` | the [presentationDefinition](https://identity.foundation/presentation-exchange/#presentation-definition) describing the required credentials.                                                                                                                        |
| `input.credentialJwts`         | a set of [verifiable credentials secured as a JWT](https://www.w3.org/TR/vc-data-model/#json-web-token)                                                                                                                                                              |
| `output.selectedCredentials`   | the expected [VC JWTs](https://identity.foundation/presentation-exchange/#presentation-submission) that satisfy the input presentation definition. If the set of VC JWTs do not satisfy the presentation definition, `output.selectedCredentials` is an empty array. |

## `validate_definition.json`
Vectors for validating whether a Presentation Definition is valid or not

| Property                       | Description                                                                                                              |
|--------------------------------|--------------------------------------------------------------------------------------------------------------------------|
| `input.presentationDefinition` | the [presentationDefinition](https://identity.foundation/presentation-exchange/#presentation-definition) to be validated |
| `errors`                       | `true` is the presentation definition is NOT valid. `false` otherwise.                                                   |

## `validate_submission.json
Vectors for validating whether a Presentation Submission is valid or not

| Property                       | Description                                                                                              |
|--------------------------------|----------------------------------------------------------------------------------------------------------|
| `input.presentationSubmission` | the [presentationSubmission](https://identity.foundation/presentation-exchange/#presentation-submission) to be validated. |
| `errors`                       | `true` is the presentation submission is NOT valid. `false` otherwise.                                   |