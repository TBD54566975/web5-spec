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
  - [Adding/Updating Vectors](#addingupdating-vectors)
  - [Feature Completeness By SDK](#feature-completeness-by-sdk)
- [Web5 SDK Features](#web5-sdk-features)
  - [Cryptographic Digital Signature Algorithms (DSA)](#cryptographic-digital-signature-algorithms-dsa)
  - [Key Management](#key-management)
  - [`did:web`](#didweb)
  - [`did:jwk`](#didjwk)
  - [`did:dht`](#diddht)
  - [`did:key`](#didkey)
  - [`did:ion`](#didion)
  - [DID Document \& Resolution Validation](#did-document--resolution-validation)
  - [W3C Verifiable Credential Data Model 1.1](#w3c-verifiable-credential-data-model-11)
  - [W3C Verifiable Credential Data Model 2.0](#w3c-verifiable-credential-data-model-20)
  - [SD-JWT / SD-JWT-VC](#sd-jwt--sd-jwt-vc)
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

- Propose API surface design _prior_ to implementation. This can be done by creating a draft PR for a respective feature and providing prototypal code or proposing the design in the comments

- Approval Required from a minimum of 2 people

> [!NOTE]
> requiring two reviewers will likely slow things down at the benefit of ensuring there's a sufficient amount of visibility on the changes being made. It also prevents bottlenecks from forming. Many people on our team should be able to speak to the features landing in our SDKs

### Release Criteria

- Each Release will have an accompanying Github Release

- Each Github Release will have an accompanying git tag that matches the version being published
- For the forseeable future, each SDK is free to publish releases at a frequency that the SDK's respective DRI sees fit

### CI / CD

Each SDK will use Github Actions for CI/CD and other automations

| Feature                       | Typescript | Kotlin | Rust | Swift |
| ----------------------------- | ---------- | ------ | ---- | ----- |
| OSS License Check             | ✅          | ✅      | ✅    | ❌     |
| Security Scanning             | ✅          | ✅      | ⛔️    | ❌     |
| Static Analysis Linting/Style | ✅          | ✅      | ✅    | ❌     |
| Running Unit Tests            | ✅          | ✅      | ✅    | ❌     |
| Publishing Tests Reports      | ✅          | ✅      | ❌    | ❌     |
| Code Coverage (CodeCov)       | ✅          | ✅      | ❌    | ❌     |
| Publishing Artifacts          | ✅          | ✅      | ❌    | ❌     |
| Release Template Checklist    | ❌          | ❌      | ❌    | ❌     |
| Automated GH Release Tag      | ❌          | ❌      | ❌    | ❌     |
| Publishing API Reference Docs | ✅          | ✅      | ❌    | ❌     |
| Publish Example Feature Usage | ✅          | ✅      | ❌    | ❌     |

> [!CAUTION]
> Security scanning via Snyk is currently not supported in Rust

- GitHub Actions should run in secured runners
  - A secure, authoritative build environment ensures software is compiled and packaged in a controlled, tamper-resistant setting.
  - This mitigates the risk of introducing vulnerabilities or malicious code during the build process, whether through external attacks or compromised internal components.
  - These runners are going to be TBD-owned and self hosted
- Ideally the above table should be represented by a "Software Catalog" with all of our SDK statuses in real time.
  - The dashboard would be consuming the data sources (GitHub, CodeCov, Snyk, Npm and other registries etc.)
  - Tools like Grafana, Backstage, or even Jenkins (weather flag) could aggregate them

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
> At a _minimum_, API reference documentation will be published to the respective sdk repository's Github Pages. e.g. `https://tbd54566975.github.io/tbdex-kt/`

---

| Language   | Comment Convention     | Docs Generator |
| ---------- | ---------------------- | -------------- |
| Typescript | TSDoc                  | API Extractor  |
| Kotlin     | KDoc                   | Dokka          |
| Rust       | Documentation comments | rustdoc        |
| Swift      | Swift Markup           | DocC           |
| Go         | tbd                    | tbd            |

> [!IMPORTANT]
> Producing API reference documentation is the responsibility of an _implementer_

### Example Feature Usage

Each SDK will **publish** example usage for _each_ implemented feature. This can either be included as a part of API reference documentation _or_ published separately

## Test Vectors

Test vectors ensure interoporability of features across SDKs and language implementations by providing common test cases with an input and expected output pair. They include both success and failure cases that can be vectorized.

This repo serves as the home for all web5 feature related vectors. They are available in the [test-vectors](./test-vectors/) directory

The `tbdex` repo houses tbdex feature related vectors. They are available in the [test-vectors](https://github.com/TBD54566975/tbdex/test-vectors) directory

The `sdk-report-runner` repo consumes the output tests for these test vectors in each repo and generates a report - [report-runner](https://github.com/TBD54566975/sdk-report-runner)

### Adding/Updating Vectors

New test vectors should follow the standard [vector structure](./test-vectors/). Vectors are automatically validated against the JSON schema via CI.

Create a PR in this repo for adding / updating web5 test vectors

### Feature Completeness By SDK

Test vectors are also used to determine feature completeness via our [test harness](./test-harness/README.md). Results of test harness runs can be found [here](https://tbd54566975.github.io/sdk-development/).

## Web5 SDK Features

### Cryptographic Digital Signature Algorithms (DSA)

| Algorithm       | Typescript | Kotlin | Rust | Swift |
| --------------- | ---------- | ------ | ---- | ----- |
| `ES256K`        | ✅          | ✅      | ✅    | ❌     |
| `EdDSA:Ed25519` | ✅          | ✅      | ✅    | ❌     |

> [!IMPORTANT]
> In-memory signing using Secp256k1 **MUST** produce k-deterministic low-s signatures. Verification **must not require** low-s signatures

### Key Management

Each SDK will implement a consistent and extensible _public interface_ for key management minimally providing concrete implementations of:

| Feature               | Typescript | Kotlin | Rust | Swift |
| --------------------- | ---------- | ------ | ---- | ----- |
| Key Manager Interface | ❌          | ✅      | ✅    | ❌     |
| In-Memory Key Manager | ❌          | ✅      | ✅    | ❌     |
| AWS KMS               | ❌          | ✅      | N/A  | N/A   |
| Device Enclave        | N/A        | ❌      | N/A  | ❌     |
| Keychain              | N/A        | ❌      | N/A  | ❌     |

Further, the key manager interface **must** be passed as an argument to _all_ public API methods that require key material. e.g.

- DID Creation
- Data Signing

> [!IMPORTANT]
> AWS KMS does **not** support `Ed25519`. At some point in the future, we can consider introducing support for key wrapping
<!-- markdownlint-disable-next-line -->

> [!IMPORTANT]
> Consumers of our SDKs should be able to provide their own `KeyManager` implementations if desired
<!-- markdownlint-disable-next-line -->

> [!NOTE]
> ⚠️ = implemented but no test vectors present

### [`did:web`](https://w3c-ccg.github.io/did-method-web/)

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| `Resolution` | ❌          | ❌      | ⚠️    | ❌     |

### [`did:jwk`](https://github.com/quartzjer/did-jwk/blob/main/spec.md)

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| `Creation`   | ❌          | ❌      | ⚠️    | ❌     |
| `Resolution` | ❌          | ❌      | ⚠️    | ❌     |

### [`did:dht`](https://tbd54566975.github.io/did-dht-method/)

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| `Creation`   | ⚠️          | ⚠️      | ❌    | ❌   |
| `Resolution` | ⚠️          | ⚠️      | ❌    | ❌   |

### [`did:key`](https://w3c-ccg.github.io/did-method-key/)

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| `Creation`   | ⚠️          | ⚠️      | ⚠️    | ❌   |
| `Resolution` | ⚠️          | ⚠️      | ⚠️    | ❌   |

> [!IMPORTANT]
> `did:key` is included because it has been implemented in both Kotlin and Typescript. I'll be creating a Github issue soon to discuss when we think it makes sense to remove ION support from both SDKs

### [`did:ion`](https://identity.foundation/sidetree/spec)

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| `Creation`   | ⚠️          | ⚠️      | ❌    | ❌   |
| `Resolution` | ⚠️          | ⚠️      | ❌    | ❌   |

> [!IMPORTANT]
> `did:ion` is included because it has been implemented in both Kotlin and Typescript. I'll be creating a Github issue soon to discuss when we think it makes sense to remove ION support from both SDKs

### [DID Document](https://www.w3.org/TR/did-core/) & [Resolution Validation](https://w3c-ccg.github.io/did-resolution/)

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| Common Error | ❌         | ❌     | ❌    | ❌   |

### [W3C Verifiable Credential Data Model 1.1](https://www.w3.org/TR/vc-data-model)

| Feature                              | Typescript | Kotlin | Rust | Swift |
| ------------------------------------ | ---------- | ------ | ---- | ----- |
| Creation                             | ✅         | ✅     | ❌   | ❌    |
| Signing as `vc-jwt`                  | ✅         | ✅     | ❌   | ❌    |
| Verification                         | ✅         | ✅     | ❌   | ❌    |
| Validation                           | ✅         | ✅     | ❌   | ❌    |
| Verifiable Presentations as `vp-jwt` | ⚠️          | ⚠️      | ❌   | ❌    |

### [W3C Verifiable Credential Data Model 2.0](https://www.w3.org/TR/vc-data-model-2.0/)

| Feature                   | Typescript | Kotlin | Rust | Swift |
| ------------------------- | ---------- | ------ | ---- | ----- |
| Creation                  | ❌         | ❌     | ❌   | ❌    |
| Signing as `vc-jose-cose` | ❌         | ❌     | ❌   | ❌    |
| Verification              | ❌         | ❌     | ❌   | ❌    |
| Validation                | ❌         | ❌     | ❌   | ❌    |

### [SD-JWT](https://datatracker.ietf.org/doc/draft-ietf-oauth-selective-disclosure-jwt/) / [SD-JWT-VC](https://datatracker.ietf.org/doc/draft-ietf-oauth-sd-jwt-vc/)

| Feature             | Typescript | Kotlin | Rust | Swift |
| ------------------- | ---------- | ------ | ---- | ----- |
| Creation            | ❌         | ❌     | ❌   | ❌    |
| Signing             | ❌         | ❌     | ❌   | ❌    |
| Verification        | ❌         | ❌     | ❌   | ❌    |
| Validation          | ❌         | ❌     | ❌   | ❌    |

### [Bitstring Status List](https://w3c.github.io/vc-bitstring-status-list/)

| Feature                           | Typescript | Kotlin | Rust | Swift |
| --------------------------------- | ---------- | ------ | ---- | ----- |
| Creation using `vc-jwt`           | ❌         | ❌     | ❌   | ❌   |
| Validation using `vc-jwt`         | ❌         | ❌     | ❌   | ❌   |

### [VC JSON Schema](https://www.w3.org/TR/vc-json-schema/)

| Feature                           | Typescript | Kotlin | Rust | Swift |
| --------------------------------- | ---------- | ------ | ---- | ----- |
| Creation `JsonSchema`             | ❌         | ❌     | ❌   | ❌   |
| Creation `JsonSchemaCredential`   | ❌         | ❌     | ❌   | ❌   |
| Validation `JsonSchema`           | ❌         | ❌     | ❌   | ❌   |
| Validation `JsonSchemaCredential` | ❌         | ❌     | ❌   | ❌   |

### [Presentation Exchange V2](https://identity.foundation/presentation-exchange/spec/v2.0.0/)

| Feature                | Typescript | Kotlin | Rust | Swift |
| ---------------------- | ---------- | ------ | ---- | ----- |
| Concrete Type          | ✅         | ✅     | ❌   | ❌    |
| Validation             | ✅         | ✅     | ❌   | ❌    |
| Credential Evaluation  | ✅         | ✅     | ❌   | ❌    |
| [Predicates](https://identity.foundation/presentation-exchange/spec/v2.0.0/#predicate-feature)                           | ✅         | ✅     | ❌   | ❌    |
| [Relational Constraints](https://identity.foundation/presentation-exchange/spec/v2.0.0/#relational-constraint-feature)   | ✅         | ❌     | ❌   | ❌    |
| [Credential Status](https://identity.foundation/presentation-exchange/spec/v2.0.0/#credential-status-constraint-feature) | ❌         | ❌     | ❌   | ❌    |
