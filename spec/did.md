# W3C Decentralized Identifiers v1.0

The [DID Core data model](https://www.w3.org/TR/did-core) provides optionality for many of its properties, as it is designed as an [abstract data model](https://www.w3.org/TR/did-core/#representations). This means that many properties can be represented in different ways while still being spec conformant, properties can take on multiple concrete types (i.e. sometimes a string, sometimes an array), and documents can be extended to include additional properties either through the [DID Specification Registry](https://www.w3.org/TR/did-spec-registries/) or via usage of [JSON-LD](https://www.w3.org/TR/json-ld11/).

This optionality can be difficult to implement consistently across languages. As a consequence, this specification defines a strict subset of the DID Core v1.0 data model represented by the following table. As a utility JSON Schemas for [DID Documents](did-document.json), [DID Resolution Metadata](did-resolution.json), and [DID Document Metadata](did-metadata.json) are provided to aid in the validation of conformant documents.

## DID Document Data Model

Following from [this data model](https://www.w3.org/TR/did-core/#core-properties).

| Property      | JSON Representation | Required | Notes          |
| ------------- | ------------------- | -------- | -------------- |
| `id`          | String              | Yes      | Must be a URI. |
| `@context`    | Array of strings    | No       | Depends on the DID method. |
| `controller`  | Array of strings    | No       | Depends on the DID method. Strings must be URIs. |
| `alsoKnownAs`        | Array of strings | No   | Depends on the DID method. Strings must be URIs. |
| `verificationMethod` | Array of [Verification Methods](#verification-method-data-model) | Yes | There must be at least one Verification Method in each DID Document. |
| `authentication`     | Array of strings | No   | String values must be fully qualified DID URIs (e.g. `did:example:abcd#key-1` over `#key-1`). |
| `assertionMethod`    | Array of strings | No   | String values must be fully qualified DID URIs (e.g. `did:example:abcd#key-1` over `#key-1`). |
| `keyAgreement`       | Array of strings | No   | String values must be fully qualified DID URIs (e.g. `did:example:abcd#key-1` over `#key-1`). |
| `capabilityInvocation` | Array of strings | No | String values must be fully qualified DID URIs (e.g. `did:example:abcd#key-1` over `#key-1`). |
| `capabilityDelegation` | Array of strings | No | String values must be fully qualified DID URIs (e.g. `did:example:abcd#key-1` over `#key-1`). |
| `service`            | Array of [Services](#service-data-model) | No | - |

**Additional Notes:**
- No [JSON-LD processing](https://www.w3.org/TR/did-core/#consumption-0) is performed.
- Each [Verification Method](https://www.w3.org/TR/did-core/#verification-methods) must have at least one [Verification Relationship](https://www.w3.org/TR/did-core/#verification-relationships).

### Verification Method Data Model

Following form [this data model](https://www.w3.org/TR/did-core/#verification-methods).

| Property      | JSON Representation | Required | Notes          |
| ------------- | ------------------- | -------- | -------------- |
| `id`          | String              | Yes      | Must be a fully qualified DID URI (e.g. `did:example:abcd#key-1`). |
| `type`        | String              | Yes      | Must be a URI. |
| `controller`  | String              | Yes      | Must be a URI. |
| `publicKeyJwk` | Object             | Yes      | Represents a [JWK](https://www.w3.org/TR/did-core/#bib-rfc7517). |

### Service Data Model

Following from [this data model](https://www.w3.org/TR/did-core/#services).

| Property          | JSON Representation | Required | Notes          |
| ----------------- | ------------------- | -------- | -------------- |
| `id`              | String              | Yes      | Must be a fully qualified DID URI (e.g. `did:example:abcd#service-1`). |
| `type`            | String              | Yes      | Must be a type defined in the [service registry](https://www.w3.org/TR/did-spec-registries/#service-types). |
| `serviceEndpoint` | Array of Strings    | Yes      | String values must be URIs. |
| `sig`             | Array of Strings    | No       | -                           |
| `enc`             | Array of Strings    | No       | -                           |

## DID Resolution Metadata Data Model

DID Resolution Metadata is _always optional_. This means that conformant implementations need not support the metadata, though it may be returned when interacting with DID resolvers.

Following from [this data model](https://www.w3.org/TR/did-core/#did-resolution-metadata).

| Property      | JSON Representation | Required | Notes          |
| ------------- | ------------------- | -------- | -------------- |
| `error`       |  String             | No       | Required if there was an error during resolution. One of the defined [error types](#did-resolution-metadata-error-types). |

### DID Resolution Metadata Error Types

Error types supported following from [this data model](https://www.w3.org/TR/did-core/#did-resolution-metadata).

| Type          | Description         |
| ------------- | ------------------- |
| `invalidDid`  | The requested DID was not valid and resolution could not proceed. |
| `notFound`    | The requested DID was not found. |
| `representationNotSupported` | The requested representation of the DID payload is not supported by the resolver. |
| `methodNotSupported` | The requested DID method is not supported by the resolver. |
| `invalidDidDocument` | The DID Document was found but did not represent a conformant document. |
| `invalidDidDocumentLength` | The size of the DID Document was not within the method's acceptable limit. |
| `internalError` | Something went wrong during DID resolution. |


## DID Document Metadata Data Model

DID Document Metadata is _always optional_. This means that conformant implementations need not support the metadata, though it may be returned when interacting with DID resolvers.

Following from [this data model](https://www.w3.org/TR/did-core/#did-document-metadata).

| Property   | JSON Representation  | Required | Notes          |
| ---------- | -------------------- | -------- | -------------- |
| `created`  | String               | No       | [XML Datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) value for when the DID was created. |
| `updated`  | String               | No       | [XML Datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) value for when the DID was last updated. |
| `deactivated` | Boolean           | No       | Required to be `true` if the DID is deactivated. |
| `nextUpdate`  | String            | No       | [XML Datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) value for when the next update of the DID. |
| `versionId`   | String            | No       | Represents the version of the last update operation. |
| `nextVersionId`| String           | No       | Represents the version of the next update operation. |
| `equivalentId` | Array of Strings | No       | A stronger form of the `alsoKnownAs` property, guaranteed by the DID method. See [this spec text](https://www.w3.org/TR/did-core/#h-note-10) for more information. |
| `canonicalId`  | String           | No       | Similar to `equivalentId`, though always a single value, never a set. See [this spec text](https://www.w3.org/TR/did-core/#dfn-canonicalid) for more information. |
