Web5 Specification v1.0
========================================

**Specification Status**: Draft

**Latest Draft**: https://github.com/TBD54566975/web5-spec

**Draft Created**: October 12, 2023

**Last Updated**: July 30, 2024

**Editors**:
~ [Frank Hinek](https://github.com/frankhinek)
~ [Gabe Cohen](https://github.com/decentralgabe)
~ [Moe Jangda](https://github.com/mistermoe)

**Contributors**:
~ [Andres Uribe](https://github.com/andresuribe87)
~ [Finn Herzfeld](https://github.com/finn-tbd)
~ [Jiyoon Koo](https://github.com/jiyoontbd)
~ [Kendall Weihe](https://github.com/KendallWeihe)
~ [Neal Roessler](https://github.com/nitro-neal)
~ [Tom Daffrun](https://github.com/tomdaffurn)

**Participate**:
~ [GitHub repo](https://github.com/TBD54566975/web5-spec)
~ [File a bug](https://github.com/TBD54566975/web5-spec/issues)
~ [Commit history](https://github.com/TBD54566975/web5-spec/commits/main)

## Abstract

Web5 is a decentralized web platform that empowers users with control over their data and identity. Built on open standards and protocols, Web5 enables the integration of decentralized identity into applications, allowing developers to create user-centric experiences while ensuring individual ownership and control over personal information. This specification selects standards from the digital identity and implements sensible constraints to enable functional, practical, and interoperable implementations of these standards. Following this ethos, the Web5 Specification defines the core components, protocols, and APIs that constitute the Web5 ecosystem, enabling interoperable and privacy-preserving decentralized applications.

## Introduction

### Background

The evolution of the World Wide Web has been marked by significant shifts in how users interact with online content and services:

1. **Web1**: Read-only, static websites
2. **Web2**: Read-write, interactive platforms
3. **Web3**: Read-write-own, blockchain-based applications
4. **Web5**: Decentralized, user-centric internet, without the need for a blockchain

Web5 addresses the limitations of previous web iterations by prioritizing user sovereignty and data portability while maintaining the ease of use and developer-friendliness of Web2. It combines the best aspects of Web2 and Web3 to create a truly decentralized and user-centric internet experience.

The primary goals of Web5 are:

1. To provide users with complete control over their digital identities
2. To enable secure, decentralized data storage and sharing
3. To simplify the development of decentralized applications
4. To promote interoperability and open standards
5. To enhance privacy and security in online interactions

### Scope

This specification covers the following key aspects of Web5 and its specifics to aid in building consistent conformant implementations:

1. Cryptographic primitives and key management
2. [[ref:Decentralized Identifiers]] (DIDs) and DID methods
3. [[ref:Verifiable Credentials]] and Presentations and accompanying functionality
4. Protocols for making use of [[ref:Decentralized Identifiers]] and [[ref:Verifiable Credentials]]

::: note
At present, [Decentralized Web Nodes](https://identity.foundation/decentralized-web-node/spec/), and other forms of decentralized storage, are out of scope of this specification.
:::

## Conformance

As well as sections marked as non-normative, all authoring guidelines, diagrams, examples, and notes in this specification are non-normative. Everything else in this specification is normative.

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in [[spec:RFC2119]].

## Terminology

This specification uses the following terms:

[[def:Decentralized Identifier, Decentralized Identifiers, DID, DIDs, DID Document, DID Documents]]
~ A [W3C specification](https://www.w3.org/TR/did-core/) describing an identifier that enables verifiable, decentralized digital identity. A DID identifier is associated with a DID Document containing keys, services, and other data. 

[[def:Verifiable Credential, Verifiable Credentials, VC, VCs]]
~ A cryptographically secure, privacy-respecting, and machine-verifiable digital credential. There are many such data models in the W3C, ISO, IETF, and other standards development organizations (SDOs).

[[def:Verifiable Presentation, Verifiable Presentations, VP, VPs]]
~ A cryptographically secure, privacy-respecting, and machine-verifiable digital credential that contains an additional signature from a _holder_ who is authorized to present the credential to a relying party.

[[def:Web5 SDK, Web5 SDKs]]
~ A software development kit that provides developers with the tools and APIs necessary to build Web5-compatible applications.

[[def:DID Resolver, DID Resolvers, DID Resolution]]
~ A DID Resolver is a software service designed to perform resolution of a [[ref:DID Document]], which is the process of obtaining a [[ref:DID Document]] for a given [[ref:DID]]. See [[ref:DID-IDENTIFIER-RESOLUTION]].

## Web5 Core Features

The Web5 SDK provides a comprehensive set of features for building decentralized protocols and applications. The following section outline the core features required to build a conformant [[ref:Web5 SDK]].

### Cryptographic Primitives

Web5 supports the following cryptographic key types for use with the noted corresponding digital signature algorithms:

| Key Type | Algorithm | Function |
|----------|-----------|----------|
| [[ref:secp256k1]] | `ES256K` [[spec:RFC8812]] | Signing and Verification |
| [[ref:Ed25519]] | `EdDSA` [[spec:RFC8032]] | Signing and Verification |

::: note
In-memory signing using [[ref:secp256k1]] ****MUST**** produce k-deterministic low-s signatures with [ECDSA](https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm) as per [[spec:RFC6979]]. Verification ****MUST NOT**** require low-s signatures.
:::

### Key Management

Web5 implementations ****MUST**** provide a consistent and extensible _public_ interface for key management, with the following minimum concrete implementations:

1. Key Manager Interface
2. In-Memory Key Manager
3. AWS KMS
4. Device Enclave (Mobile)
5. Keychain (Mobile)

::: todo
Provide detailed specifications for each key management implementation, including APIs and usage guidelines.
:::

Further, the key manager interface ****MUST**** be passed as an argument to all public API methods that require key material such as:
* DID Creation
* Data Signing

::: note
[AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/asymmetric-key-specs.html#key-spec-ecc) does not support [[ref:Ed25519]]. At some point in the future, we can consider introducing support for key wrapping.
:::

Consumers of conformant [[ref:Web5 SDKs]] ****SHOULD**** be able to provide their own `KeyManager` impelementations, if desired.

### JSON Schema

[[ref:JSON-SCHEMA]] support is ****REQUIRED**** for multiple features throughout [[ref:Web5 SDKs]], including, but not limited to support with [[ref:Verifiable Credentials]], [VC JSON Schemas](#vc-json-schema), and for use in validating the data models this specification has defined.

Conformant implementations of [[ref:Web5 SDKs]] ****MUST**** support at least [[ref:JSON-SCHEMA-DRAFT-7]] and [[ref:JSON-SCHEMA-2020-12]].

:::todo
Elaborate on required features for JSON Schema support such as constructing and validating JSON Schemas. Provide guidance on caching and offline resolution.
:::

### Decentralized Identifiers (DIDs)

Decentralized Identifiers (DIDs) are a core component of Web5, providing a foundation for self-sovereign identity and decentralized authentication, authorization, and discovery of services. Web5 implements a subset of the [[ref:DID-CORE]] specification to ensure consistency and interoperability across implementations. More detail can be found in the [section on the DID Document Data Model](#did-document-data-model).

DIDs in Web5 are globally unique identifiers that enable verifiable, decentralized digital identity. They are designed to be independent of centralized registries, identity providers, or certificate authorities. Key features of [[ref:DIDs] in Web5 include:

1. **Decentralization**: [[ref:DIDs]] can be created and managed without needing to rely on centralized authorities.
2. **Control**: The controller has full say over their [[ref:DID]] and associated [[ref:DID Document]].
3. **Persistence**: [[ref:DIDs]] are persistent identifiers that do not require the continued operation of an underlying organization.
4. **Resolvability**: [[ref:DIDs]] can be resolved to retrieve associated [[ref:DID Documents]] containing cryptographic material, service endpoints, and other data.
4. **Cryptographic Verifiability**: [[ref:DIDs]] enable cryptographic verification of claims and interactions.

#### Supported DID Methods

Web5 supports the following DID methods:

| Method | Creation | Resolution | Note |
|--------|----------|------------|------|
| [[ref:did:web]] | ❌ | ✅ | - |
| [[ref:did:jwk]] | ✅ | ✅ | - |
| [[ref:did:dht]] | ✅ | ✅ | This is the "default" method for [[ref:Web5 SDKs]]. |
| [[ref:did:key]] | ⚠️ | ⚠️ | This method has been implemented in both Kotlin and TypeScript, with no plans for support in other languages. |

#### DID Features

**DID Operations**:

Conformant [[ref:Web5 SDKs]] ****MUST**** support the following DID operations:

1. **Creation**: Generate new DIDs using supported methods.
2. **Resolution**: Resolve DIDs to retrieve their associated DID Documents.
3. **Update**: Modify DID Documents for methods that support updates.
3. **Deactivation**: Deactivate DIDs when they are no longer needed or compromised.

**DID Document Management**:

Conformant [[ref:Web5 SDKs]] ****MUST**** provide functionality to:

1. Create and validate DID Documents according to the specified data model.
2. Manage verification methods and services associated with DIDs.
3. Handle DID Document metadata and resolution metadata.

::: todo
Provide more detailed information on the implementation and use of each supported DID method.
:::

### Verifiable Credentials (VCs)

Verifiable Credentials (VCs) are another cornerstone of Web5, providing a standard way to express and verify claims about entities. Web5 implements a subset of the [[ref:VC-DATA-MODEL-1.1]] specification to ensure consistency and interoperability.

[[ref:VCs]] in Web5 are cryptographically secure, privacy-respecting, and machine-verifiable digital credentials. They enable the assertion and verification of claims about subjects in a decentralized manner. Key features of VCs in Web5 include:

1. **Cryptographic Security**: VCs are digitally signed, ensuring their integrity and authenticity.
2. **Privacy Preservation**: VCs enable user-control with features like selective disclosure.
3. **Standardization**: VCs follow standardized data models, promoting interoperability across industries.
4. **Flexibility**: VCs can represent a wide range of credentials and attributes.
5. **Verifiability**: Claims in VCs can be independently verified by relying parties.

::: note
Future versions of this specification may adopt other models of [[ref:verifiable credentials]], such as [ISO mDLs](https://www.iso.org/standard/69084.html), [SD-JWT-based Verifiable Credentials](https://datatracker.ietf.org/doc/draft-ietf-oauth-sd-jwt-vc/), or [W3C Verifiable Credentials v2.0](https://www.w3.org/TR/vc-data-model-2.0/).
:::

#### Supported Credential Types

Web5 primarily supports the W3C Verifiable Credentials Data Model v1.1 [[ref:VC-DATA-MODEL-1.1]], with the following features:

1. **Basic Credentials**: Standard Verifiable Credentials containing claims about a subject.
2. **[[ref:Verifiable Presentations]]**: Holder signed VCs that can be shared with verifiers.
3. **JSON Schema**: VCs that use [[ref:VC-JSON-SCHEMA]] for claim validation.
4. **Status List Credentials**: VCs that support revocation and suspension using [[ref:STATUS-LIST-2021]].

#### VC Features

Conformant [[ref:Web5 SDKs]] ****MUST**** support the following [[ref:Verifiable Credential]] operations:

1. **Issuance**: Create and sign [[ref:Verifiable Credentials]] using the `vc-jwt` format.
2. **Verification**: Verify the integrity, authenticity, and status of [[ref:Verifiable Credentials]] as `vc-jwt`.
3. **Presentation Creation**: Generate [[ref:Verifiable Presentations]] from one or more [[ref:VCs]] as `vp-jwt`.
4. **Presentation Verification**: Verify the integrity and authenticity of [[ref:Verifiable Presentations]] as `vp-jwt`.
5. **Status**: Manage the status of issued credentials using [[ref:STATUS-LIST-2021]].

:::todo
[](https://github.com/TBD54566975/web5-spec/issues/12) Specify credential validation in more detail.
:::

#### VC Formats

Web5 supports the following credential formats:

1. **JWT**: Verifiable Credentials and Presentations secured as JSON Web Tokens (JWTs) [[spec:RFC7519]] using the `vc-jwt` and `vp-jwt` types.
2. **JSON-LD**: Support for [[ref:JSON-LD]] context in VCs, although [[ref:JSON-LD]] processing is not performed.

#### Credential Status

Conformant [[ref:Web5 SDKs]] ****MUST**** support Credential Status as specified by [[ref:STATUS-LIST-2021]] for usage with the W3C Verifiable Credential Data Model 1.1 [[ref:VC-DATA-MODEL-1.1]].

The following operations ****MUST**** be supported:

1. **Issuance**: Create and sign [[ref:Verifiable Credentials]] representing [[ref:STATUS-LIST-2021]] VCs as `vc-jwt`.
2. **Verification**: Verify the integrity and authenticity of [[ref:STATUS-LIST-2021]] VCs as `vc-jwt`.
3. **Status Setting**: Support setting multiple statuses including `Revoked` and `Suspended`.
4. **Status Checking**: Functionality to determine whether a credential has a given status.

:::note
A new version of this status list, under development as [[ref:BITSTRING-STATUS-LIST]], is under consideration for adoption in a future version of this specification.
:::

:::todo
Add more detail on expected features and data model constraints.
:::

#### Credential Schema

Conformant [[ref:Web5 SDKs]] ****MUST**** support [[ref:VC-JSON-SCHEMA]] for usage with the W3C Verifiable Credential Data Model 1.1 [[ref:VC-DATA-MODEL-1.1]].

The following operations ****MUST**** be supported:
1. **Creation**: The ability to construct a [[ref:VC-JSON-SCHEMA]] of type `JsonSchema`.
2. **Resolution**: The ability to resolve a JSON Schema from a web resource given its `id`.
3. **Validation**: Functionality to validate a given [[ref:Verifiable Credential]] against a [[ref:VC-JSON-SCHEMA]].

:::todo
Add more detail on expected features and data model constraints.
:::

### Presentation Exchange

Web5 incorporates Presentation Exchange v2.0 [[ref:PRESENTATION-EXCHANGE]] as a crucial component for facilitating the exchange of Verifiable Credentials between parties. This feature enables standardized request and submission of credentials. Presentation Exchange provides a mechanism for one party to request specific credentials or claims from another party, and for the responding party to present those credentials in a standardized format. This enables dynamic, privacy-preserving, and minimal disclosure of information.

#### Presentation Exchnage Features

Conformant [[ref:Web5 SDKs]] ****MUST**** support the following [[ref:PRESENTATION-EXCHANGE]] features:

1. **Data Model**: Implementation of the Presentation Exchange data model, including Presentation Definition, Presentation Submission, and related structures.
2. **Validation**: Ability to validate Presentation Definitions and Presentation Submissions against the Presentation Exchange schema.
3. **Credential Evaluation**: Support for evaluating credentials across multiple formats (`vc-jwt` and `vp-jwt` using [[ref:VC-DATA-MODEL-1.1]].
4. **Predicates**: Support for predicate-based claims, allowing for filtering of claims against [[ref:JSON-SCHEMA]] values.
5. **Relational Constraints**: Ability to express and evaluate constraints such as `subjet_is_issuer`, `same_subject`, or `is_holder`.
6. **Credential Status**: Support for including and evaluating credential status (see: [Credential Status](#credential-status)) information as part of the presentation exchange process.

#### Presentation Exchange Operations

Conformant [[ref:Web5 SDKs]] ****MUST**** support the following operations in processing Presentation Exchange data objects:

1. **Presentation Request Creation**: Generate Presentation Definitions specifying the required credentials and claims, signed as a `vc-jwt`.
2. **Presentation Request Parsing**: Verify the signature, validate the data, and interpret received Presentation Definitions.
3. **Presentation Submission Generation**: Create Presentation Submissions that satisfy the requirements of a Presentation Definition. Support signing and verification using `vc-jwt`.
4. **Presentation Submission Verification**: Verify that a Presentation Submission meets the criteria specified in a Presentation Definition.
	* Evaluate predicates.
	* Evaluate relational constraints.
	* Evaluate credential status.

::: note
Implementers should refer to the [[ref:PRESENTATION-EXCHANGE]] specification for detailed requirements and guidelines on implementing these features.
:::

::: todo
Add specific examples of how Presentation Exchange is used in common Web5 scenarios, such as authentication workflows or sharing of user attributes.
:::

## Data Models

This section provides guidance for implementers of [[ref:Web5 SDKs]] [core features](#web5-sdk-core-features).

### DID Document Data Model

[[ref:DID-CORE]] provides optionality for many of its properties, as it is designed as an [abstract data model](https://www.w3.org/TR/did-core/#representations). This means that many properties can be represented in different ways while still being spec conformant, properties can take on multiple concrete types (i.e. sometimes a string, sometimes an array), and documents can be extended to include additional properties either through the [DID Specification Registry](https://www.w3.org/TR/did-spec-registries/) or via usage of [JSON-LD](https://www.w3.org/TR/json-ld11/).

This optionality can be difficult to implement consistently across languages. As a consequence, this specification defines a strict subset of the DID Core v1.0 data model represented by the following table. As a utility JSON Schemas for [DID Documents](did-document.json), [DID Resolution Metadata](did-resolution.json), and [DID Document Metadata](did-metadata.json) are provided to aid in the validation of conformant documents.

Following from [this data model](https://www.w3.org/TR/did-core/#core-properties).

| Property      | JSON Representation | Required | Guidance       |
| ------------- | ------------------- | -------- | -------------- |
| `id`          | String              | Yes      | ****MUST**** be a URI. |
| `@context`    | Array of strings    | No       | Depends on the DID method. |
| `controller`  | Array of strings    | No       | Depends on the DID method. Strings ****MUST**** be URIs. |
| `alsoKnownAs`        | Array of strings | No   | Depends on the DID method. Strings ****MUST**** be URIs. |
| `verificationMethod` | Array of [Verification Methods](#verification-method-data-model) | Yes | There ****MUST**** be at least one Verification Method in each DID Document. |
| `authentication`     | Array of strings | No   | String values ****MUST**** be fully qualified DID URIs (e.g. `did:example:abcd#key-1` over `#key-1`). |
| `assertionMethod`    | Array of strings | No   | String values ****MUST**** be fully qualified DID URIs (e.g. `did:example:abcd#key-1` over `#key-1`). |
| `keyAgreement`       | Array of strings | No   | String values ****MUST**** be fully qualified DID URIs (e.g. `did:example:abcd#key-1` over `#key-1`). |
| `capabilityInvocation` | Array of strings | No | String values ****MUST**** be fully qualified DID URIs (e.g. `did:example:abcd#key-1` over `#key-1`). |
| `capabilityDelegation` | Array of strings | No | String values ****MUST**** be fully qualified DID URIs (e.g. `did:example:abcd#key-1` over `#key-1`). |
| `service`            | Array of [Services](#service-data-model) | No | - |

**Additional Guidance**:
- [JSON-LD processing](https://www.w3.org/TR/did-core/#consumption-0) ****MUST NOT**** be performed.
- Each [Verification Method](https://www.w3.org/TR/did-core/#verification-methods) ****MUST**** have at least one [Verification Relationship](https://www.w3.org/TR/did-core/#verification-relationships).

#### Verification Method Data Model

Following form [this data model](https://www.w3.org/TR/did-core/#verification-methods).

| Property      | JSON Representation | Required | Guidance       |
| ------------- | ------------------- | -------- | -------------- |
| `id`          | String              | Yes      | ****MUST**** be a fully qualified DID URI (e.g. `did:example:abcd#key-1`). |
| `type`        | String              | Yes      | ****MUST**** be a URI. |
| `controller`  | String              | Yes      | ****MUST**** be a URI. |
| `publicKeyJwk` | Object             | Yes      | Represents a JWK [[spec:RFC7517]]. |

#### Service Data Model

Following from [this data model](https://www.w3.org/TR/did-core/#services).

| Property          | JSON Representation | Required | Guidance       |
| ----------------- | ------------------- | -------- | -------------- |
| `id`              | String              | Yes      | ****MUST**** be a fully qualified DID URI (e.g. `did:example:abcd#service-1`). |
| `type`            | String              | Yes      | ****MUST**** be a type defined in the [service registry](https://www.w3.org/TR/did-spec-registries/#service-types). |
| `serviceEndpoint` | Array of Strings    | Yes      | String values ****MUST**** be URIs. |
| `sig`             | Array of Strings    | No       | -                           |
| `enc`             | Array of Strings    | No       | -                           |

### DID Resolution Metadata Data Model

DID Resolution Metadata ****MAY**** be present. This means that conformant implementations are ****NOT REQUIRED**** to support the metadata, though it ****MAY**** be returned when interacting with [[ref:DID Resolvers]].

Following from [this data model](https://www.w3.org/TR/did-core/#did-resolution-metadata).

| Property      | JSON Representation | Required | Notes          |
| ------------- | ------------------- | -------- | -------------- |
| `error`       |  String             | No       | Required if there was an error during resolution. ****MUST**** be one of the defined [error types](#did-resolution-metadata-error-types). |

#### DID Resolution Metadata Error Types

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

### DID Document Metadata Data Model

DID Document Metadata is _always optional_. This means that conformant implementations ****MAY NOT**** support the metadata, though it ****MAY**** be returned when interacting with [[ref:DID Resolvers]].

Following from [this data model](https://www.w3.org/TR/did-core/#did-document-metadata).

| Property   | JSON Representation  | Required | Guidance       |
| ---------- | -------------------- | -------- | -------------- |
| `created`  | String               | No       | [[ref:XML Datetime]] value for when the DID was created. |
| `updated`  | String               | No       | [[ref:XML Datetime]] value for when the DID was last updated. |
| `deactivated` | Boolean           | No       | Required to be `true` if the DID is deactivated. |
| `nextUpdate`  | String            | No       | [[ref:XML Datetime]] value for when the next update of the DID. |
| `versionId`   | String            | No       | Represents the version of the last update operation. |
| `nextVersionId`| String           | No       | Represents the version of the next update operation. |
| `equivalentId` | Array of Strings | No       | A stronger form of the `alsoKnownAs` property, guaranteed by the DID method. See [this spec text](https://www.w3.org/TR/did-core/#h-note-10) for more information. |
| `canonicalId`  | String           | No       | Similar to `equivalentId`, though always a single value, never a set. See [this spec text](https://www.w3.org/TR/did-core/#dfn-canonicalid) for more information. |


### Verifiable Credentials v1.1 Data Model

The [[ref:VC Data Model]] provides optionality for many of its properties. This means they can take on multiple concrete types, for example an [`issuer` property](https://www.w3.org/TR/vc-data-model/#issuer) can either be repreesnted as a JSON string representing a URI or a JSON object that must contain an `id` property. Additionally, the data model follows an [open world](https://www.w3.org/TR/vc-data-model/#extensibility) model for extensibility via the usage of [[ref:JSON-LD]]. As a consequence, conformant Verifiable Credentials may contain properties that are not defined by the specification itself, but by [JSON-LD Contexts](https://www.w3.org/TR/json-ld11/#the-context).

This optionality can be difficult to implement consistently across languages. As a consequence, this specification defines a strict subset of the VC Data Model v1.1 that supports the [plain JSON syntax](https://www.w3.org/TR/vc-data-model/#json), represented by the following table. As a utility JSON Schemas for [Verifiable Credentials](vc-11.json) and [Verifiable Presentations](vp-11.json) are provided to aid in the validation of conformant documents.

Following from [this data model](https://www.w3.org/TR/vc-data-model/#basic-concepts).

| Property      | JSON Representation | Required | Guidance       |
| ------------- | ------------------- | -------- | -------------- |
| `@context`    | Array of strings    | Yes      | Contexts defining the meaning of terms within the credential. ****MUST**** include at least `"https://www.w3.org/2018/credentials/v1"`. |
| `id`          | String              | Yes      | A URI representing a unique identifier for the credential. ****RECOMMENDED**** to be of form `urn:uuid:3978344f-8596-4c3a-a978-8fcaba3903c5`. |
| `type`        | Array of strings    | Yes      | Type(s) of the credential. Must include `VerifiableCredential`. |
| `issuer`      | String OR Object    | Yes      | ****RECOMMENDED**** to be a string; a DID representing a unique identifier for the entity that issued the credential. We also need to support the case where `issuer` is a JSON Object with an `id` propertery (following prior guidance) and a `name` property representing the Issuer's name. |
| `issuanceDate`| String              | Yes      | [XML Datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) value for when the credential was issued. |
| `expirationDate` | String           | No       | [XML Datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) value after which the credential is no longer valid. |
| `credentialSubject` | Object        | Yes      | Data about the subject of the credential. Can be any JSON object. |
| `credentialSubject.id` | String     | Yes      | A [[ref:DID]] representing a unique identifier for whom the credential's claims are made. |
| `credentialStatus` | Object defined by [Credential Status](#credential-status) | No | Only to be used with [[ref:STATUS-LIST-2021]]. |
| `credentialSchema` | Object defined by [Credential Schema](#credential-schema) | No | Recommended. Only to be used with the type [`JsonSchema`](https://www.w3c.org/TR/vc-json-schema/#jsonschema). |
| `evidence`    | Array of objects    | No       | An array of JSON objects as per [Evidence](https://www.w3.org/TR/vc-data-model/#evidence). |

**Additional Guidance**:
- The `credentialSubject` property can be any JSON object. It is ****RECOMMENDED**** that this object is defined by an associated `credentialSchema`.
- No [JSON-LD processing](https://www.w3.org/TR/vc-data-model/#json-ld) is performed.
- Embedded proofs, using the `proof` property ****MUST NOT**** be present. JWTs with the `proof` property present ****MUST NOT**** be processed.
- The `type` property ****MUST**** always contain `VerifiableCredential` and ****MAY**** also contain the URI(s) of a [[ref:JSON Schema]], if one is used for the credential.
- Verifiable Credentials ****MUST NOT** support multiple credential subjects.
- Verifiable Credentials ****MUST**** be secured as JWTs according to the [rules laid out in the specification](https://www.w3.org/TR/vc-data-model/#json-web-token).
- XML Datetime values ****MAY**** be represented by conforming to [[ref:ISO8601]] or [[spec:RFC3339]] formats, as they are subsets of XML Datetime.
- For the `evidence` property no further implementation is needed until we are able to specify an evidence [type](https://www.w3.org/TR/vc-data-model/#dfn-type), such as those provided by [this registry](https://w3c.github.io/vc-specs-dir/#evidence) [[ref:VC-SPECIFICATIONS-DIRECTORY]].

#### Credential Status Data Model

Following from [this data model](https://www.w3.org/community/reports/credentials/CG-FINAL-vc-status-list-2021-20230102/).

**StatusList2021Entry**

| Property        | JSON Representation | Required | Guidance       |
| --------------- | ------------------- | -------- | -------------- |
| `id`            | String              | Yes      | A URL which uniquely identifies the status of the associated verifiable credential. |
| `type`          | String              | Yes      | ****MUST**** be set to `StatusList2021Entry`. |
| `statusPurpose` | String              | Yes      | Describes the type of status the object represents (e.g. `revocation` or `suspension`). |
| `statusListIndex` | String            | Yes      | An integer >= 0 expressed as a string that identifies the bit position of the status of the associated verifiable credential. |
| `statusListCredential` | String       | Yes      | A URL which uniquely identifies a verifiable credential whose type is `StatusList2021Credential`. |


:::note
When representing Status List Credentials, as opposed to including a status in another VC, the Status List Credential must [follow the guidance here](https://www.w3.org/community/reports/credentials/CG-FINAL-vc-status-list-2021-20230102/#statuslist2021credential).
:::

#### Credential Schema Data Model

Following from [this data model](https://www.w3.org/TR/vc-json-schema/#jsonschema).

| Property        | JSON Representation | Required | Notes          |
| --------------- | ------------------- | -------- | -------------- |
| `id`            | String              | Yes      | A URL which uniquely identifies the JSON Schema for the associated Verifiable Credential. |
| `type`          | String              | Yes      | Must be set to `JsonSchema`. |

:::note
Although [the referenced spec](https://w3c.org/TR/vc-json-schema/) is designed for v2 of the VC Data Model, we apply it to v1.1 as a standard means to implement the `credentialSchema`.
:::

### Verifiable Presentation v1.1 Data Model

Following from [this guidance](https://www.w3.org/TR/vc-data-model/#presentations-0), which extends on the [Verifiable Credentials Data Model](#w3c-verifiable-credentials-v11) above.

| Property      | JSON Representation | Required | Notes          |
| ------------- | ------------------- | -------- | -------------- |
| `@context`    | Array of strings    | Yes      | Contexts defining the meaning of terms within the presentation. Must include at least `"https://www.w3.org/2018/credentials/v1"`. |
| `id`          | String              | Yes      | A URI representing a unique identifier for the presentation. Recommended to be of form `urn:uuid:3978344f-8596-4c3a-a978-8fcaba3903c5`. |
| `type`        | Array of strings    | Yes      | Type(s) of the presentation. Must include `VerifiablePresentation`. |
| `holder`      | String              | Yes      | A DID representing a unique identifier for the entity that created the presentation. |
| `issuanceDate`| String              | Yes      | [XML Datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) value for when the presentation was created. |
| `expirationDate` | String           | No       | [XML Datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) value after which the presentation is no longer valid. |
| `verifiableCredential` | Array of strings | Yes    | An array with at least one value, containing the JWT representation of [Verifiable Credential](#verifiable-credential-data-model) objects. |

**Additional Guidance**:
- No [JSON-LD processing](https://www.w3.org/TR/vc-data-model/#json-ld) is performed.
- Embedded proofs, using the `proof` property ****MUST NOT**** be present.
- The `type` property ****MUST**** always contain `VerifiablePresentation` and ****MAY**** also contain the URI(s) of a [[ref:JSON Schema]], if one is used for the presentation.
- Verifiable Presentations ****MUST**** be secured as JWTs [[spec:RFC7519]] according to the [rules laid out in the specification](https://www.w3.org/TR/vc-data-model/#json-web-token).
- [[ref:XML Datetime]] values ****MAY*** be represented by conforming to [[ref:ISO8601]] or [[spec:RFC3339]] formats, as they are subsets of XML Datetime.

## Additional Features

::: todo
Note additional features of Web5 SDKs, if any.
:::

## Security Considerations

::: todo
Discuss security considerations specific to Web5, including:
- Key management best practices
- Secure communication protocols
- Threat models and mitigation strategies
- Cryptographic algorithm choices and their implications
:::

## Privacy Considerations

::: todo
Discuss privacy considerations specific to Web5, including:
- Data minimization principles
- User consent and control mechanisms
- Pseudonymity and unlinkability features
- Compliance with privacy regulations (e.g., GDPR, CCPA)
:::

## Implementation Guidelines

Find the latest [Projects Health Dashboard here](https://developer.tbd.website/open-source/projects-dashboard/).

::: todo
Provide guidelines for implementing Web5 applications, including:
- Best practices for using Web5 SDKs
- Common patterns and architectures
- Performance optimization techniques
- Error handling and recovery strategies
:::

## Versioning and Compatibility

| Component | Version | Status | Notes |
|-----------|---------|--------|-------|
| Web5 Specification | 1.0 | Draft | This document |
| DID Core | 1.0 | [[def:DID-CORE]] | With adjustments as specified in [this section](#did-document-data-model). |
| Verifiable Credentials Data Model | 1.1 | [[ref:VC-DATA-MODEL-1.1]] | With adjustments as specified in [this section](#verifiable-credentials-v11-data-model). |
| Presentation Exchange | 2.1.1 | [[ref:PRESENTATION-EXCHANGE]] | With adjustments as specified in [this section](#presentation-exchange). | 
| Cryptographic Algorithms | - | - | As specified in [the following section](#cryptographic-primitives). |

::: todo
Provide guidance on versioning strategy and backward compatibility considerations for future Web5 specification updates.
:::

## References

[[def:DID-CORE]]
~ [Decentralized Identifiers (DIDs) v1.0](https://www.w3.org/TR/did-core/). W3C Recommendation, 19 July 2022. M. Sporny, A. Guy, M. Sabadello, D. Reed. [W3C](https://www.w3.org/).

[[def:DID-IDENTIFIER-RESOLUTION]]
~ [Decentralized Identifier Resolution (DID Resolution) v0.3](https://w3c.github.io/did-resolution/). Draft Community Group Report, 24 July 2024. M. Sabadello, D. Zagidulin. [W3C](https://www.w3.org/).

[[def:VC-DATA-MODEL-1.1, VC Data Model]]
~ [Verifiable Credentials Data Model v1.1](https://www.w3.org/TR/vc-data-model/). W3C Recommendation, 03 March 2022. M. Sprony, G. Noble, D. Longley, D. Burnett, B. Zundel, K. D. Hartog. [W3C](https://www.w3.org/).

[[def:VC-JSON-SCHEMA]]
~ [Verifiable Credentials JSON Schema Specification](https://www.w3.org/TR/vc-json-schema/). G. Cohen, M. Prorock, A. Uribe; 18 December 2023. [W3C](https://www.w3.org/).

[[def:JSON-SCHEMA]]
~ [JSON Schema](https://json-schema.org/specification.html). A Media Type for Describing JSON Documents. OpenJS Foundation.

[[def:JSON-SCHEMA-DRAFT-7]]
~ [JSON Schema Draft 7](https://json-schema.org/draft-07/json-schema-release-notes.html). OpenJS Foundation.

[[def:JSON-SCHEMA-2020-12]]
~ [JSON Schema 2020-12](https://json-schema.org/draft/2020-12/release-notes.html). OpenJS Foundation.

[[def:STATUS-LIST-2021]]
~ [Status List 2021](https://www.w3.org/community/reports/credentials/CG-FINAL-vc-status-list-2021-20230102/). Privacy-preserving status information for Verifiable Credentials. M. Sporny, D. Longley, O. Steele, M. Prorock, M. Alkhraishi; 02 January 2023. [W3C CCG](https://w3c-ccg.github.io).

[[def:BITSTRING-STATUS-LIST]]
~ [Bitstring Status List v1.0](https://www.w3.org/TR/vc-bitstring-status-list/). Privacy-preserving status information for Verifiable Credentials. D. Longley, M. Sporny, O. Steele. [W3C](https://www.w3.org/).

[[def:VC-SPECIFICATIONS-DIRECTORY]]
~ [VC Specifications Directory](https://www.w3.org/TR/vc-specs-dir/). M. Sporny. W3C Group Note. 01 June 2024. [W3C](https://www.w3.org/).

[[def:DID JWK, did:jwk]]
~ [did:jwk](https://github.com/quartzjer/did-jwk/blob/main/spec.md). did:jwk is a deterministic transformation of a
JWK into a DID Document. J. Miller.

[[def:DID Web, did:web]]
~ [did:web Method Specification](https://w3c-ccg.github.io/did-method-web/). C. Gribneau, M. Prorock, O. Steele,
O. Terbu, M. Xu, D. Zagidulin; 06 May 2023. [W3C](https://www.w3.org/).

[[def:DID DHT, did:dht]]
~ [did:dht](https://did-dht.com/). G. Cohen, D. Buchner. Implementers Draft. July 25, 2024. [TBD](https://tbd.website/).

[[def:DID Key, did:key]]
~ [The did:key Method v0.7](https://w3c-ccg.github.io/did-method-key/). A DID Method for Static Cryptographic Keys.
D. Longley, D. Zagidulin, M. Sporny. [W3C CCG](https://w3c-ccg.github.io/).

[[def:JSON-LD]]
~ [JSON-LD 1.1]. A JSON-based Serialization for Linked Data. W3C Recommendation. G. Kellogg, P. A. Champin, D. Longley. 16 July 2020. [W3C](https://www.w3.org/).

[[def:PRESENTATION-EXCHANGE]]
~ [Presentation Exchange 2.1.1](https://identity.foundation/presentation-exchange/spec/v2.1.1/). D. Buchner, B. Zundel, M. Riedel, K. Hamilton Duffy; DIF Ratified Specification. June 30, 2024. [DIF](https://www.identity.foundation).

[[def:Ed25519]]
~ [Ed25519](https://ed25519.cr.yp.to/). D. J. Bernstein, N. Duif, T. Lange, P. Schwabe, B.-Y. Yang; 26 September 2011.
[ed25519.cr.yp.to](https://ed25519.cr.yp.to/).

[[def:secp256k1]]
~ [secp256k1](https://www.secg.org/sec2-v2.pdf). Certicom Research. January 27, 2010. [SECG](https://www.secg.org/).

[[def:ISO8601]]
~ [ISO8601](https://www.iso.org/iso-8601-date-and-time-format.html). Date and time format. Edition 1, 2019. [ISO](https://www.iso.org).

[[def:XML Datetime]]
~ [W3C XML Schema Definition Language (XSD) 1.1 Part 2: Datatypes](https://www.w3.org/TR/xmlschema11-2/#dateTime). D. Peterson, S. Gao, A. Malhotra, C. M. Sperberg-McQueen, H. S. Thompson. W3C Recommendation. 5 April 2012. [W3C](https://www.w3.org/).

[[spec]]