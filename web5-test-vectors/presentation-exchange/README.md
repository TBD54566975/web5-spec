# Presentation Exchange Test Vectors

## CreatePresentationFromCredentials

Input and output for a full presentation exchange test vectors are available [here](./wa-license.json)

### Input

the value of `input` is a an object with `presentationDefinition` and the corresponding `credentialJwt`

| Property                | Description                                                                                                                                                                                |
| ----------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `presentationDefinition`           | the input [presentationDefinition](https://identity.foundation/presentation-exchange/#presentation-definition)  showing the requirements used for this getting an example WA license               |
| `credentialJwt`   | the input [credentialJwt](https://www.w3.org/TR/did-core/#dfn-diddocumentmetadata) that corresponds to the presentationDefinition to fulfill it and do a full presentation exchange

### Output

the value of `output` is an object that contains the following properties

| Property                | Description                                                                                                                                                                                |
| ----------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `presentationSubmission`           | the expected [presentationSubmission](https://www.w3.org/TR/did-core/#dfn-diddocument) when the `inputs` are processed by `createPresentationFromCredentials`.             |
