# W3C Verifiable Credentials v1.1

The VC Data Model provides optionality for many of its properties. This means they can take on multiple concrete types, for example an [`issuer` property](https://www.w3.org/TR/vc-data-model/#issuer) can either be repreesnted as a JSON string representing a URI or a JSON object that must contain an `id` property. Additionally, the data model follows an [open world](https://www.w3.org/TR/vc-data-model/#extensibility) model for extensibility via the usage of [JSON-LD](https://www.w3.org/TR/json-ld11/). As a consequence, conformant Verifiable Credentials may contain properties that are not defined by the specification itself, but by [JSON-LD Contexts](https://www.w3.org/TR/json-ld11/#the-context).

This optionality can be difficult to implement consistently across languages. As a consequence, this specification defines a strict subset of the VC Data Model v1.1 that supports the [plain JSON syntax](https://www.w3.org/TR/vc-data-model/#json), represented by the following table. As a utility JSON Schemas for [Verifiable Credentials](spec/vc-11.json) and [Verifiable Presentations](spec/vp-11.json) are provided to aid in the validation of conformant documents.

## Verifiable Credentials Data Model

## 