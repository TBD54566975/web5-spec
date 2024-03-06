# Web5-Spec <!-- omit in toc -->

- [Purpose](#purpose)
- [Requirements](#requirements)
  - [Feature Tracking](#feature-tracking)
  - [Feature Acceptance](#feature-acceptance)
  - [Work Prioritization](#work-prioritization)
  - [Implementation Criteria](#implementation-criteria)
  - [Review Criteria](#review-criteria)
  - [Release Criteria](#release-criteria)
  - [CI / CD](#ci--cd)
  - [Publishing Artifacts](#publishing-artifacts)
  - [Publishing API Reference Documentation](#publishing-api-reference-documentation)
  - [Example Feature Usage](#example-feature-usage)
- [Test Vectors](#test-vectors)
  - [Adding & Updating Vectors](#adding--updating-test-vectors)
  - [Feature Completeness By SDK](#feature-completeness-by-sdk)
- [Web5 SDK Features](#web5-sdk-features)
  - [Cryptographic Digital Signature Algorithms (DSA)](#cryptographic-digital-signature-algorithms-dsa)
  - [Key Management](#key-management)
  - [Decentralized Identifiers](#decentralized-identifiers)
    - [DID Documents & Resolution](#did-documents--did-resolution)
    - [DID Methods](#did-methods)
  - [Verifiable Credentials](#verifiable-credentials)
    - [W3C Verifiable Credentials v1.1](#w3c-verifiable-credential-data-model-11)
    - [W3C Verifiable Credentials v2.0](#w3c-verifiable-credential-data-model-20)
    - [SD-JWT](#sd-jwt)
    - [Status List 2021](#status-list-2021)
    - [Bitstring Status List](#bitstring-status-list)
    - [VC JSON Schema](#vc-json-schema)
    - [Presentation Exchange V2](#presentation-exchange-v2)

## Purpose

This repo sets forth the development process, requirements, and the desired feature for the following SDKs:

- [web5-js](https://github.com/TBD54566975/web5-js)
- [web5-kt](https://github.com/TBD54566975/web5-kt)
- [web5-go](https://github.com/TBD54566975/web5-go)
- [web5-rs](https://github.com/TBD54566975/web5-rs)
- [web5-swift](https://github.com/TBD54566975/web5-swift)

## Requirements

### Feature Tracking

Feature tracking will take place entirely on GitHub through Github Issues and the [SDK development Github project](https://github.com/orgs/TBD54566975/projects/29).

### Feature Acceptance

Proposing new features that impact all SDKs will occur by creating a Github Issue in this repo. The Issue should include motivation or rationale in addition to any relevant reading material. New features will be discussed and decided upon during weekly syncs

> [!IMPORTANT]
> Language agnostic test vectors **must** be produced _prior_ to commencing implementation beyond the first SDK

Test vectors are in the test-vector directory. More info on test vectors can be found there.

### Work Prioritization

Prioritization of features or specific work items will be reflected during weekly sync meetings.

### Implementation Criteria

An individual SDK will consider a feature implemented once the following requirements have been met:

- Produce/update API reference documentation for any/all methods added to the _public_ API surface. documenting private API methods is optional
- Produce/update relevant example usage
- Test coverage must be provided for _all_ public API methods.
- Tests that consume shared test vectors must all be passing

### Review Criteria

- Propose API surface design _prior_ to implementation. This can be done by opening a feature request Issue and fostering a discussion around the proposed design. This can be followed by a corresponding draft PR and providing prototypal code.

- Approval Required from a minimum of 2 people.

> [!NOTE]
> Requiring two reviewers will likely slow things down at the benefit of ensuring there's a sufficient amount of visibility on the changes being made. It also prevents bottlenecks from forming. Many people on our team should be able to speak to the features landing in our SDKs.

### Release Criteria

- Each Release will have an accompanying Github Release

- Each Github Release will have an accompanying git tag that matches the version being published
- For the forseeable future, each SDK is free to publish releases at a frequency that the SDK's respective DRI sees fit

### CI / CD

Each SDK will use GitHub Actions for CI/CD and other automations.

Find the latest [Projects Health Dashboard here](https://developer.tbd.website/open-source/projects-dashboard/). 

### Publishing Artifacts

Each SDK will be published to the most widely adopted registry/repository for the respective language

| SDK        | Repository          |
| ---------- | ------------------- |
| Typescript | npm                 |
| Kotlin     | maven central       |
| Rust       | crates.io           |
| Swift      | swift package index |
| Go         | tbd                 |

### Publishing API Reference Documentation

Each SDK will auto generate API reference documentation using the respective language's commenting convention and doc gen tooling

---

> [!IMPORTANT]
> At a _minimum_, API reference documentation will be published to the respective sdk repository's Github Pages. e.g. `https://tbd54566975.github.io/tbdex-kt/`.

---

| Language   | Comment Convention     | Docs Generator |
| ---------- | ---------------------- | -------------- |
| Typescript | TSDoc                  | API Extractor  |
| Kotlin     | KDoc                   | Dokka          |
| Rust       | Documentation comments | rustdoc        |
| Swift      | Swift Markup           | DocC           |
| Go         | tbd                    | tbd            |

> [!IMPORTANT]
> Producing API reference documentation is the responsibility of an _implementer_.

### Example Feature Usage

Each SDK will **publish** example usage for _each_ implemented feature. This can either be included as a part of API reference documentation _or_ published separately.

## Test Vectors

Test vectors ensure interoporability of features across SDKs and language implementations by providing common test cases with an input and expected output pair. They include both success and failure cases that can be vectorized.

This repo serves as the home for all web5 feature related vectors. They are available in the [test-vectors](./test-vectors/) directory.

The `tbdex` repo houses tbdex feature related vectors. They are available in the [test-vectors](https://github.com/TBD54566975/tbdex/test-vectors) directory.

The `sdk-report-runner` repo consumes the output tests for these test vectors in each repo and generates a report - [report-runner](https://github.com/TBD54566975/sdk-report-runner).

### Adding & Updating Test Vectors

New test vectors should follow the standard [vector structure](./test-vectors/). Vectors are automatically validated against their [JSON Schema](https://json-schema.org/) via CI.

Create a PR in this repo for adding / updating Web5 test vectors.

### Feature Completeness By SDK

## Web5 SDK Features

Test vectors are used to determine feature completeness via our [test harness](./test-harness/README.md). Results of test harness runs can be found [here](https://tbd54566975.github.io/sdk-development/).

### Cryptographic Digital Signature Algorithms (DSA)

| Key Type                                          | Algorithm                                                             | Type                     |
| ------------------------------------------------- | --------------------------------------------------------------------- | ------------------------ |
| [secp256k1](https://en.bitcoin.it/wiki/Secp256k1) | [`ES256K`](https://datatracker.ietf.org/doc/html/rfc8812#section-3.1) | Signing and Verification |
| [Ed25519](https://ed25519.cr.yp.to/)              | [`EdDSA`](https://datatracker.ietf.org/doc/html/rfc8032)              | Signing and Verification |

> [!IMPORTANT]
> In-memory signing using secp256k1 **MUST** produce k-deterministic low-s signatures with [ECDSA](https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm) as per [RFC6979](https://www.rfc-editor.org/rfc/rfc6979). Verification **must not require** low-s signatures.

### Key Management

Each SDK will implement a consistent and extensible _public interface_ for key management minimally providing concrete implementations of:

| Feature               |
| --------------------- |
| Key Manager Interface | 
| In-Memory Key Manager |
| AWS KMS               |
| Device Enclave        |
| Keychain              |

Further, the key manager interface **must** be passed as an argument to _all_ public API methods that require key material. e.g.

- DID Creation
- Data Signing

> [!IMPORTANT]
> AWS KMS does **not** support `Ed25519`. At some point in the future, we can consider introducing support for key wrapping
<!-- markdownlint-disable-next-line -->

> [!IMPORTANT]
> Consumers of our SDKs should be able to provide their own `KeyManager` implementations if desired
<!-- markdownlint-disable-next-line -->

### Decentralized Identifiers

#### [DID Documents](https://www.w3.org/TR/did-core/) & [DID Resolution](https://w3c-ccg.github.io/did-resolution/)

Independent of DID Methods, we support DID Documents according to [DID Core v1.0](https://www.w3.org/TR/2022/REC-did-core-20220719/) with the adjustments specified in [this corresponding document](spec/did.md).

#### DID Methods

| Method                                                 | Creation   | Resolution | Note |
| ------------------------------------------------------ | ---------- | ---------- | ---- |
| [`did:web`](https://w3c-ccg.github.io/did-method-web/) | ❌         | ✅         | -   |
| [`did:jwk`](https://github.com/quartzjer/did-jwk/blob/main/spec.md) | ✅         | ✅  | - |
| [`did:dht`](https://did-dht.com)                       | ✅         | ✅         | This is our default method. |
| [`did:key`](https://w3c-ccg.github.io/did-method-key/) | ⚠️         | ⚠️         | Has been implemented in both Kotlin and Typescript, no plans for support in other languages. |
| [`did:ion`](https://identity.foundation/sidetree/spec) | ⚠️         | ⚠️         | Support for `did:ion` has been deprecated. |

### Verifiable Credentials

#### [W3C Verifiable Credential Data Model 1.1](https://www.w3.org/TR/vc-data-model)

We support Verifiable Credentials according to the [Verifiable Credentials Data Model 1.1](https://www.w3.org/TR/vc-data-model) with the adjustments specified in [this corresponding document](spec/vc.md).

| Feature                                                          |
| ---------------------------------------------------------------- |
| Data model                                                       |
| Validation (data model, JSON Schema, status)                     |
| Signing and verification of Verifiable Credentials as `vc-jwt`   |
| Signing and verification of Verifiable Presentations as `vp-jwt` |

#### [W3C Verifiable Credential Data Model 2.0](https://www.w3.org/TR/vc-data-model-2.0/)

| Feature                                                                  |
| ------------------------------------------------------------------------ |
| Data model                                                               |
| Validation (data model, JSON Schema, status)                             |
| Signing and verification of Verifiable Credentials with `vc-jose-cose`.  |
| Signing and verification of Verifiable Presentations with `vc-jose-cose` |

#### [SD-JWT](https://datatracker.ietf.org/doc/draft-ietf-oauth-selective-disclosure-jwt/)

| Feature                  |
| ------------------------ |
| Data model               | 
| Signing and verification | 

#### [Status List 2021](https://www.w3.org/community/reports/credentials/CG-FINAL-vc-status-list-2021-20230102/)

For usage with the [W3C Verifiable Credential Data Model 1.1](https://www.w3.org/TR/vc-data-model).

| Feature                                | 
| -------------------------------------- | 
| Data model                             |
| Status checking                        |
| Status setting                         |
| Signing and verification with `vc-jwt` |

#### [Bitstring Status List](https://www.w3.org/TR/vc-bitstring-status-list/)

For usage with the [W3C Verifiable Credential Data Model 2.0](https://www.w3.org/TR/vc-data-model-2.0/).

| Feature                                      | 
| -------------------------------------------- | 
| Data model                                   |
| Status checking                              |
| Status setting                               |
| Signing and verification with `vc-jose-cose` |


#### [VC JSON Schema](https://www.w3.org/TR/vc-json-schema/)

For usage with both the [W3C Verifiable Credential Data Model 1.1](https://www.w3.org/TR/vc-data-model) and the [W3C Verifiable Credential Data Model 2.0](https://www.w3.org/TR/vc-data-model-2.0/).

| Feature                                 |
| --------------------------------------- | 
| Creation and validation of `JsonSchema` |
| Creation and validation of `JsonSchemaCredential` using `vc-jwt` with the VCDM v1.1 |
| Creation and validation of `JsonSchemaCredential` using `vc-jose-cose` with the VCDM v2.0 |

#### [Presentation Exchange v2.0](https://identity.foundation/presentation-exchange/spec/v2.0.0/)

| Feature                | 
| ---------------------- | 
| Data model             | 
| Validation             |
| Credential Evaluation using the VCDM v1.1 (both `vc-jwt` and `vp-jwt`), the VCDM v2.0 using `vc-jose-cose`, and SD-JWT   |
| [Predicates](https://identity.foundation/presentation-exchange/spec/v2.0.0/#predicate-feature)                           |
| [Relational Constraints](https://identity.foundation/presentation-exchange/spec/v2.0.0/#relational-constraint-feature)   |
| [Credential Status](https://identity.foundation/presentation-exchange/spec/v2.0.0/#credential-status-constraint-feature) |