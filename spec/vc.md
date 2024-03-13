# W3C Verifiable Credentials v1.1

The VC Data Model provides optionality for many of its properties. This means they can take on multiple concrete types, for example an [`issuer` property](https://www.w3.org/TR/vc-data-model/#issuer) can either be repreesnted as a JSON string representing a URI or a JSON object that must contain an `id` property. Additionally, the data model follows an [open world](https://www.w3.org/TR/vc-data-model/#extensibility) model for extensibility via the usage of [JSON-LD](https://www.w3.org/TR/json-ld11/). As a consequence, conformant Verifiable Credentials may contain properties that are not defined by the specification itself, but by [JSON-LD Contexts](https://www.w3.org/TR/json-ld11/#the-context).

This optionality can be difficult to implement consistently across languages. As a consequence, this specification defines a strict subset of the VC Data Model v1.1 that supports the [plain JSON syntax](https://www.w3.org/TR/vc-data-model/#json), represented by the following table. As a utility JSON Schemas for [Verifiable Credentials](vc-11.json) and [Verifiable Presentations](vp-11.json) are provided to aid in the validation of conformant documents.

## Verifiable Credential Data Model

Following from [this data model](https://www.w3.org/TR/vc-data-model/#basic-concepts).

| Property      | JSON Representation | Required | Notes          |
| ------------- | ------------------- | -------- | -------------- |
| `@context`    | Array of strings    | Yes      | Contexts defining the meaning of terms within the credential. Must include at least `"https://www.w3.org/2018/credentials/v1"`. |
| `id`          | String              | Yes      | A URI representing a unique identifier for the credential. Recommended to be of form `urn:uuid:3978344f-8596-4c3a-a978-8fcaba3903c5`. |
| `type`        | Array of strings    | Yes      | Type(s) of the credential. Must include `VerifiableCredential`. |
| `issuer`      | String              | Yes      | A DID representing a unique identifier for the entity that issued the credential. |
| `issuanceDate`| String              | Yes      | [XML Datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) value for when the credential was issued. |
| `expirationDate` | String           | No       | [XML Datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) value after which the credential is no longer valid. |
| `credentialSubject` | Object        | Yes      | Data about the subject of the credential. Can be any JSON object. |
| `credentialSubject.id` | String     | Yes      | A DID representing a unique identifier for whom the credential's claims are made. |
| `credentialStatus` | Object defined by [Credential Status](#credential-status) | No | Only to be used with [Status List 2021](https://www.w3.org/community/reports/credentials/CG-FINAL-vc-status-list-2021-20230102/). |
| `credentialSchema` | Object defined by [Credential Schema](#credential-schema) | No | Recommended. Only to be used with the type [`JsonSchema`](https://w3c.github.io/vc-json-schema/#jsonschema). |
| `evidence`    | Object              | No       | Any JSON object as per [Evidence](https://www.w3.org/TR/vc-data-model/#evidence). |

**Additional Notes:**
- The `credentialSubject` property can be any JSON object. It is recommended that this object is defined by an associated `credentialSchema`.
- No [JSON-LD processing](https://www.w3.org/TR/vc-data-model/#json-ld) is performed.
- Embedded proofs, using the `proof` property must not be present.
- The `type` property must always contain `VerifiableCredential` but may also contain the URI(s) of a JSON Schema, if one is used for the credential.
- We do not support multiple credential subjects.
- Verifiable Credentials must be secured as JWTs according to the [rules laid out in the specification](https://www.w3.org/TR/vc-data-model/#json-web-token).

### Credential Status Data Model

Following from [this data model](https://www.w3.org/community/reports/credentials/CG-FINAL-vc-status-list-2021-20230102/).

**StatusList2021Entry**

| Property        | JSON Representation | Required | Notes          |
| --------------- | ------------------- | -------- | -------------- |
| `id`            | String              | Yes      | A URL which uniquely identifies the status of the associated verifiable credential. |
| `type`          | String              | Yes      | Must be set to `StatusList2021Entry`. |
| `statusPurpose` | String              | Yes      | Describes the type of status the object represents (e.g. `revocation` or `suspension`). |
| `statusListIndex` | String            | Yes      | An integer >= 0 expressed as a string that identifies the bit position of the status of the associated verifiable credential. |
| `statusListCredential` | String       | Yes      | A URL which uniquely identifies a verifiable credential whose type is `StatusList2021Credential`. |

**Additional Notes:**
- When representing Status List Credentials, as opposed to including a status in another VC, the Status List Credential must [follow the guidance here](https://www.w3.org/community/reports/credentials/CG-FINAL-vc-status-list-2021-20230102/#statuslist2021credential).

### Credential Schema Data Model

Following from [this data model](https://w3c.github.io/vc-json-schema/#jsonschema).

| Property        | JSON Representation | Required | Notes          |
| --------------- | ------------------- | -------- | -------------- |
| `id`            | String              | Yes      | A URL which uniquely identifies the JSON Schema for the associated Verifiable Credential. |
| `type`          | String              | Yes      | Must be set to `JsonSchema`. |

**Additional Notes:**
- Although [the referenced spec](https://w3c.github.io/vc-json-schema/) is designed for v2 of the VC Data Model, we apply it to v1.1 as a standard means to implement the `credentialSchema`.

## Verifiable Presentation Data Model

Following from [this guidance](https://www.w3.org/TR/vc-data-model/#presentations-0), which extends on the data model above.

| Property      | JSON Representation | Required | Notes          |
| ------------- | ------------------- | -------- | -------------- |
| `@context`    | Array of strings    | Yes      | Contexts defining the meaning of terms within the presentation. Must include at least `"https://www.w3.org/2018/credentials/v1"`. |
| `id`          | String              | Yes      | A URI representing a unique identifier for the presentation. Recommended to be of form `urn:uuid:3978344f-8596-4c3a-a978-8fcaba3903c5`. |
| `type`        | Array of strings    | Yes      | Type(s) of the presentation. Must include `VerifiablePresentation`. |
| `holder`      | String              | Yes      | A DID representing a unique identifier for the entity that created the presentation. |
| `issuanceDate`| String              | Yes      | [XML Datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) value for when the presentation was created. |
| `expirationDate` | String           | No       | [XML Datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) value after which the presentation is no longer valid. |
| `verifiableCredential` | Array of strings | Yes    | An array with at least one value, containing the JWT representation of [Verifiable Credential](#verifiable-credential-data-model) objects. |

**Additional Notes:**
- No [JSON-LD processing](https://www.w3.org/TR/vc-data-model/#json-ld) is performed.
- Embedded proofs, using the `proof` property must not be present.
- The `type` property may only contain `VerifiablePresentation` or the URI of a JSON Schema, if one is used for the presentation.
- Verifiable Presentations must be secured as JWTs according to the [rules laid out in the specification](https://www.w3.org/TR/vc-data-model/#json-web-token).
