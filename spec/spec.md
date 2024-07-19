Web5 Specification v1.0
========================================

**Specification Status**: Draft

**Latest Draft**: https://TODO

**Draft Created**: October 12, 2023

**Last Updated**: July 19, 2024

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

Web5 is a decentralized web platform that empowers users with control over their data and identity. Built on open standards and protocols, Web5 enables the integration of decentralized identity into applications, allowing developers to create user-centric experiences while ensuring individual ownership and control over personal information. This specification selects standards from the digital identity and implemenets sensible constraints to enable functional, practical, and interoperable implementations of these standards. Following this ethos, the Web5 Specification defines the core components, protocols, and APIs that constitute the Web5 ecosystem, enabling interoperable and privacy-preserving decentralized applications.

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
2. Decentralized Identifiers (DIDs) and DID methods
3. Verifiable Credentials and Presentations and accompanying functionality
4. Protocols for making use of Decentralized Identifiers and Verifiable Credentials

::: note
At present, [Decentralized Web Nodes](https://identity.foundation/decentralized-web-node/spec/), and other forms of decentralized storage, are out of scope of this specification.
:::

## Conformance

As well as sections marked as non-normative, all authoring guidelines, diagrams, examples, and notes in this specification are non-normative. Everything else in this specification is normative.

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in [[spec:RFC2119]].

## Terminology

This specification uses the following terms:

[[def:Decentralized Identifier, Decentralized Identifier, DID, DIDs, DID Document, DID Documents]]
~ A [W3C specification](https://www.w3.org/TR/did-core/) describing an identifier that enables verifiable, decentralized digital identity. A DID identifier is associated with a DID Document containing keys, services, and other data. 

[[def:Verifiable Credential, Verifiable Credentials, VC, VCs]]
~ A cryptographically secure, privacy-respecting, and machine-verifiable digital credential. There are many such data models in the W3C, ISO, IETF, and other standards development organizations (SDOs).

[[def:Web5 SDK]]
~ A software development kit that provides developers with the tools and APIs necessary to build Web5-compatible applications.

::: todo
Expand the terminology section with additional Web5-specific terms and concepts.
:::

## Web5 SDK Features

The Web5 SDK provides a comprehensive set of features for building decentralized protocols and applications. This section outlines the core features and their implementation status across different programming languages.

### Decentralized Identifiers (DIDs)

Web5 uses [[def:Decentralized Identifiers]] (DIDs) as the foundation for user identity and authentication. This specification adopts a subset of the [[ref:DID-CORE]] specification with some adjustments to ensure consistency across implementations.

#### DID Document Data Model

The following table defines the properties of a DID Document in Web5:

| Property | JSON Representation | Required | Notes |
|----------|---------------------|----------|-------|
| `id` | String | Yes | Must be a URI. |
| `@context` | Array of strings | No | Depends on the DID method. |
| `controller` | Array of strings | No | Depends on the DID method. Strings must be URIs. |
| `alsoKnownAs` | Array of strings | No | Depends on the DID method. Strings must be URIs. |
| `verificationMethod` | Array of [Verification Methods](#52-verification-method-data-model) | Yes | There must be at least one Verification Method in each DID Document. |
| `authentication` | Array of strings | No | String values must be fully qualified DID URIs (e.g., `did:example:abcd#key-1` over `#key-1`). |
| `assertionMethod` | Array of strings | No | String values must be fully qualified DID URIs. |
| `keyAgreement` | Array of strings | No | String values must be fully qualified DID URIs. |
| `capabilityInvocation` | Array of strings | No | String values must be fully qualified DID URIs. |
| `capabilityDelegation` | Array of strings | No | String values must be fully qualified DID URIs. |
| `service` | Array of [Services](#53-service-data-model) | No | - |

#### Verification Method Data Model

| Property | JSON Representation | Required | Notes |
|----------|---------------------|----------|-------|
| `id` | String | Yes | Must be a fully qualified DID URI (e.g., `did:example:abcd#key-1`). |
| `type` | String | Yes | Must be a URI. |
| `controller` | String | Yes | Must be a URI. |
| `publicKeyJwk` | Object | Yes | Represents a [[RFC7517]] JWK. |

#### Service Data Model

| Property | JSON Representation | Required | Notes |
|----------|---------------------|----------|-------|
| `id` | String | Yes | Must be a fully qualified DID URI (e.g., `did:example:abcd#service-1`). |
| `type` | String | Yes | Must be a type defined in the [[DID-SPEC-REGISTRIES]] service registry. |
| `serviceEndpoint` | Array of Strings | Yes | String values must be URIs. |
| `sig` | Array of Strings | No | - |
| `enc` | Array of Strings | No | - |

### DID Methods

Web5 supports the following DID methods:

| Method | Creation | Resolution | Note |
|--------|----------|------------|------|
| [`did:web`](https://w3c-ccg.github.io/did-method-web/) | ❌ | ✅ | - |
| [`did:jwk`](https://github.com/quartzjer/did-jwk/blob/main/spec.md) | ✅ | ✅ | - |
| [`did:dht`](https://did-dht.com) | ✅ | ✅ | This is our default method. |
| [`did:key`](https://w3c-ccg.github.io/did-method-key/) | ⚠️ | ⚠️ | Has been implemented in both Kotlin and TypeScript, no plans for support in other languages. |
| [`did:ion`](https://identity.foundation/sidetree/spec) | ⚠️ | ⚠️ | Support for `did:ion` has been deprecated. |

::: todo
Provide more detailed information on the implementation and use of each supported DID method.
:::

## Verifiable Credentials

Web5 implements a subset of the [[ref:VC-DATA-MODEL-1.1]] specification to ensure consistency across implementations. This section defines the data models for Verifiable Credentials and Verifiable Presentations in Web5.

### Verifiable Credential Data Model

| Property | JSON Representation | Required | Notes |
|----------|---------------------|----------|-------|
| `@context` | Array of strings | Yes | Contexts defining the meaning of terms within the credential. Must include at least `"https://www.w3.org/2018/credentials/v1"`. |
| `id` | String | Yes | A URI representing a unique identifier for the credential. Recommended to be of form `urn:uuid:3978344f-8596-4c3a-a978-8fcaba3903c5`. |
| `type` | Array of strings | Yes | Type(s) of the credential. Must include `VerifiableCredential`. |
| `issuer` | String OR Object | Yes | Recommended to be a string; a DID representing a unique identifier for the entity that issued the credential. We also need to support the case where `issuer` is a JSON Object with an `id` property and a `name` property representing the Issuer's name. |
| `issuanceDate` | String | Yes | [[XMLSCHEMA11-2]] DateTime value for when the credential was issued. |
| `expirationDate` | String | No | [[XMLSCHEMA11-2]] DateTime value after which the credential is no longer valid. |
| `credentialSubject` | Object | Yes | Data about the subject of the credential. Can be any JSON object. |
| `credentialSubject.id` | String | Yes | A DID representing a unique identifier for whom the credential's claims are made. |
| `credentialStatus` | Object | No | Only to be used with [[STATUS-LIST-2021]]. |
| `credentialSchema` | Object | No | Recommended. Only to be used with the type `JsonSchema` [[VC-JSON-SCHEMA]]. |
| `evidence` | Array of objects | No | An array of JSON objects as per [Evidence](https://www.w3.org/TR/vc-data-model/#evidence). |

::: todo
Provide examples of Verifiable Credentials in Web5 and guidance on their usage.
:::

### Verifiable Presentation Data Model

| Property | JSON Representation | Required | Notes |
|----------|---------------------|----------|-------|
| `@context` | Array of strings | Yes | Contexts defining the meaning of terms within the presentation. Must include at least `"https://www.w3.org/2018/credentials/v1"`. |
| `id` | String | Yes | A URI representing a unique identifier for the presentation. Recommended to be of form `urn:uuid:3978344f-8596-4c3a-a978-8fcaba3903c5`. |
| `type` | Array of strings | Yes | Type(s) of the presentation. Must include `VerifiablePresentation`. |
| `holder` | String | Yes | A DID representing a unique identifier for the entity that created the presentation. |
| `issuanceDate` | String | Yes | [[XMLSCHEMA11-2]] DateTime value for when the presentation was created. |
| `expirationDate` | String | No | [[XMLSCHEMA11-2]] DateTime value after which the presentation is no longer valid. |
| `verifiableCredential` | Array of strings | Yes | An array with at least one value, containing the JWT representation of Verifiable Credential objects. |

::: todo
Provide examples of Verifiable Presentations in Web5 and guidance on their usage.
:::

### JSON Schema

### Credential Status

## Cryptographic Algorithms

Web5 supports the following cryptographic digital signature algorithms:

| Key Type | Algorithm | Type |
|----------|-----------|------|
| [secp256k1](https://en.bitcoin.it/wiki/Secp256k1) | [`ES256K`](https://datatracker.ietf.org/doc/html/rfc8812#section-3.1) | Signing and Verification |
| [Ed25519](https://ed25519.cr.yp.to/) | [`EdDSA`](https://datatracker.ietf.org/doc/html/rfc8032) | Signing and Verification |

::: note
In-memory signing using secp256k1 **MUST** produce k-deterministic low-s signatures with [ECDSA](https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm) as per [[spec:RFC6979]]. Verification **must not require** low-s signatures.
:::

## Key Management

Web5 implementations must provide a consistent and extensible public interface for key management, with the following minimum concrete implementations:

1. Key Manager Interface
2. In-Memory Key Manager
3. AWS KMS
4. Device Enclave (Mobile)
5. Keychain (Mobile)

::: todo
Provide detailed specifications for each key management implementation, including APIs and usage guidelines.
:::

## Additional Features

::: todo
Provide detailed specifications for the following features:
- SD-JWT
- Status List 2021
- Bitstring Status List
- VC JSON Schema
- Presentation Exchange V2
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
| DID Core | 1.0 | [[def:DID-CORE]] | With adjustments as specified in [the following section](#decentralized-identifiers-dids.) |
| Verifiable Credentials Data Model | 1.1 | [[ref:VC-DATA-MODEL-1.0]] | With adjustments as specified in [the following section](#verifiable-credentials). |
| Cryptographic Algorithms | - | - | As specified in [the following section](#cryptographic-algorithms). |

::: todo
Provide guidance on versioning strategy and backward compatibility considerations for future Web5 specification updates.
:::

## References

[[def:DID-CORE]]
~ [Decentralized Identifiers (DIDs) v1.0](https://www.w3.org/TR/did-core/). W3C Recommendation, 19 July 2022. M. Sporny, A. Guy, M. Sabadello, D. Reed.

[[def:VC-DATA-MODEL-1.1]]
~ [Verifiable Credentials Data Model v1.1](https://www.w3.org/TR/vc-data-model/). W3C Recommendation, 03 March 2022. M. Sprony, G. Noble, D. Longley, D. Burnett, B. Zundel, K. D. Hartog. [W3C](https://www.w3.org/).

[[def:VC-JOSE-COSE]]
~ [Securing Verifiable Credentials using JOSE and COSE](https://www.w3.org/TR/vc-jose-cose/). O. Steele, M. Jones,
M. Prorock, G. Cohen; 25 April 2024. [W3C](https://www.w3.org/).

[[def:DID Web, did:web]]
~ [did:web Method Specification](https://w3c-ccg.github.io/did-method-web/). C. Gribneau, M. Prorock, O. Steele,
O. Terbu, M. Xu, D. Zagidulin; 06 May 2023. [W3C](https://www.w3.org/).

[[spec]]