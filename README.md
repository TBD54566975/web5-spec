# SDK Development <!-- omit in toc -->

- [Purpose](#purpose)
- [Known Unknowns](#known-unknowns)
- [Requirements](#requirements)
  - [Feature Tracking](#feature-tracking)
    - [Labels](#labels)
    - [Milestones](#milestones)
  - [Feature Acceptance](#feature-acceptance)
  - [Work Prioritization](#work-prioritization)
  - [Implementation Criteria](#implementation-criteria)
  - [Review Criteria](#review-criteria)
  - [Release Criteria](#release-criteria)
  - [CI / CD](#ci--cd)
  - [Publishing Artifacts](#publishing-artifacts)
  - [Publishing API Reference Documentation](#publishing-api-reference-documentation)
  - [Example Feature Usage](#example-feature-usage)
- [Features](#features)
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
  - [Presentation Definition V2](#presentation-definition-v2)
  - [tbDEX Message](#tbdex-message)
  - [tbDEX Resource](#tbdex-resource)
  - [tbDEX Offering Resource](#tbdex-offering-resource)
  - [tbDEX RFQ Message](#tbdex-rfq-message)
  - [tbDEX Quote Message](#tbdex-quote-message)
  - [tbDEX Order Message](#tbdex-order-message)
  - [tbDEX Order-Status Message](#tbdex-order-status-message)
  - [tbDEX Close Message](#tbdex-close-message)


# Purpose

This repo sets forth the development process, requirements, and desired feature set for the following SDKs:
* [tbdex-js](https://github.com/TBD54566975/tbdex-js)
* [tbdex-kt](https://github.com/TBD54566975/tbdex-kt)
* [tbdex-rs](https://github.com/TBD54566975/tbdex-rs)
* [web5-js](https://github.com/TBD54566975/web5-js)
* [web5-kt](https://github.com/TBD54566975/web5-kt)
* [web5-rs](https://github.com/TBD54566975/web5-rs)


Github Issues and PRs created in this repo address topics that impact / span all SDKs we have in development. See section on [Feature Acceptance](#feature-acceptance) for information on submitting proposals for new features across all of our SDKs (e.g. implement Presentation Exchange v9)

# Known Unknowns
* We don't yet have a cemented path for supporting Android. It could either be through our pre-existing Kotlin SDK or bindings exposed via Rust
* We're not yet in a place where we can confidently state that a rust core can support a broad surface area of target languages. We're confident in pursuing rust development for the purposes of surfacing swift bindings.

# Requirements

## Feature Tracking
Feature tracking will take place entirely on GitHub through Github Issues and the [SDK development Github project](https://github.com/orgs/TBD54566975/projects/29).

---

Work items that impact all SDKs should be created as issues in _this_ Github Repo. e.g. "produce test vectors for `did:web`

---

Work items that are specific to an individual SDK should be created as an issue in the respective sdk's Github Repo

> [!IMPORTANT]
> Each individual SDK's DRI will be responsible for creating all of the relevant github issues in their respective repo by using the tables below. After all issues are created for a given feature, create a PR for this repo to delete the relevant table.

---

Work item progress is tracked using the `Status` attribute on a GH issue. This should automatically be reflected on the kanban view on the GH project

---

### Labels

> [!IMPORTANT]
> A label should be created for each feature, in each respective repo.

The following labels should exist in all relevant repos

| Label                | Color Hex | Description                                |
| -------------------- | --------- | ------------------------------------------ |
| `dsa`                | `#7FDBFF` | Cryptographic Digital Signature Algorithms |
| `key-mgmt`           | `#0074D9` | Key Management                             |
| `did:web`            | `#2ECC40` | did:web                                    |
| `did:jwk`            | `#FFDC00` | did:jwk                                    |
| `did:dht`            | `#FF851B` | did:dht                                    |
| `did:key`            | `#F012BE` | did:key                                    |
| `did:ion`            | `#B10DC9` | did:ion                                    |
| `did-doc-validation` | `#3D9970` | DID Document & Resolution Validation       |
| `w3c-vc-dm-1.1`      | `#39CCCC` | W3C Verifiable Credential Data Model 1.1   |
| `w3c-vc-dm-2.0`      | `#01FF70` | W3C Verifiable Credential Data Model 2.0   |
| `sd-jwt`             | `#85144B` | SD-JWT / SD-JWT-VC                         |
| `pd-v2`              | `#F9A602` | Presentation Definition V2                 |
| `tbdex-message`      | `#70DB93` | tbDEX Message                              |
| `tbdex-resource`     | `#5B2C6F` | tbDEX Resource                             |
| `tbdex-offering`     | `#E59866` | tbDEX Offering Resource                    |
| `tbdex-rfq`          | `#1F618D` | tbDEX RFQ Message                          |
| `tbdex-quote`        | `#186A3B` | tbDEX Quote Message                        |
| `tbdex-order`        | `#28B463` | tbDEX Order Message                        |
| `tbdex-orderstatus`  | `#D68910` | tbDEX Order-Status Message                 |
| `tbdex-close`        | `#34495E` | tbDEX Close Message                        |



> [!NOTE]
> This list will change over time as features are added or removed

---

### Milestones
Github Repo Milestones will be used to track work for specific codenamed projects (e.g. Eagle, Pigeon, Black Swan, etc.).

The following milestones should exist in all relevant repos

| Milestone |
| --------- |
| `Eagle`   |
| `ABC`     |

> [!NOTE]
> This list will change over time as projects are added

## Feature Acceptance
Proposing new features that impact all SDKs will occur by creating a Github Issue in this repo. The Issue should include motivation or rationale in addition to any relevant reading material. New features will be discussed and decided upon during weekly syncs

> [!IMPORTANT]
> Language agnostic test vectors **must** be produced _prior_ to commencing implementation beyond the first SDK


## Work Prioritization
Prioritization of features or specific work items will be reflected during weekly sync meetings.

## Implementation Criteria
An individual SDK will consider a feature implemented once the following requirements have been met:
* Produce/update API reference documentation for any/all methods added to the _public_ API surface. documenting private API methods is optional
* Produce/update relevant example usage
* Test coverage must be provided for _all_ public API methods.
* Tests that consume shared test vectors must all be passing


## Review Criteria
* Propose API surface design _prior_ to implementation. This can be done by creating a draft PR for a respective feature and providing prototypal code or proposing the design in the comments
* Approval Required from a minimum of 2 people

> [!NOTE]
> requiring two reviewers will likely slow things down at the benefit of ensuring there's a sufficient amount of visibility on the changes being made. It also prevents bottlenecks from forming. Many people on our team should be able to speak to the features landing in our SDKs


## Release Criteria
* Each Release will have an accompanying Github Release
* Each Github Release will have an accompanying git tag that matches the version being published
* For the forseeable future, each SDK is free to publish releases at a frequency that the SDK's respective DRI sees fit

## CI / CD
Each SDK will use Github Actions for CI/CD and other automations

| Feature                       | Typescript | Kotlin | Rust | Swift |
| ----------------------------- | ---------- | ------ | ---- | ----- |
| OSS License Check             | ❌          | ❌      | ❌    | ❌     |
| Security Scanning             | ❌          | ❌      | ❌    | ❌     |
| Static Analysis Linting/Style | ❌          | ❌      | ❌    | ❌     |
| Running Unit Tests            | ✅          | ✅      | ❌    | ❌     |
| Publishing Tests Reports      | ❌          | ❌      | ❌    | ❌     |
| Code Coverage (CodeCov)       | ❌          | ❌      | ❌    | ❌     |
| Publishing Artifacts          | ❓          | ❌      | ❌    | ❌     |
| Release Template Checklist    | ❌          | ❌      | ❌    | ❌     |
| Automated GH Release Tag      | ❌          | ❌      | ❌    | ❌     |
| Publishing API Reference Docs | ❌          | ✅      | ❌    | ❌     |
| Publish Example Feature Usage | ❌          | ✅      | ❌    | ❌     |

* GitHub Actions should run in secured runners
* Ideally the above table should be represented by a "Software Catalog" with all of our SDK statuses in real time.
  * The dashboard would be consuming the data sources (GitHub, CodeCov, Snyk, Npm and other registries etc.)
  * Tools like Grafana, Backstage, or even Jenkins (weather flag) could aggregate them

## Publishing Artifacts
Each SDK will be published to the most widely adopted registry/repository for the respective language


| SDK        | Repository    |
| ---------- | ------------- |
| Typescript | npm                 |
| Kotlin     | maven central       |
| Rust       | crates              |
| Swift      | swift package index |


## Publishing API Reference Documentation
Each SDK will auto generate API reference documentation using the respective language's commenting convention and doc gen tooling 

---

> [!IMPORTANT]
> At a _minimum_, API reference documentation will be published to the respective sdk repository's Github Pages. e.g. `https://tbd54566975.github.io/tbdex-kt/`

---

| Language   | Comment Convention | Docs Generator |
| ---------- | ------------------ | -------------- |
| Typescript | TSDoc              | API Extractor  |
| Kotlin     | KDoc               | Dokka          |
| Rust       | ?                  | ?              |
| Swift      | ?                  | ?              |


> [!IMPORTANT] 
> Producing API reference documentation is the responsibility of an _implementer_

## Example Feature Usage
Each SDK will **publish** example usage for _each_ implemented feature. This can either be included as a part of API reference documentation _or_ published separately

# Features

## Cryptographic Digital Signature Algorithms (DSA)
| Algorithm       | Typescript | Kotlin | Rust | Swift |
| --------------- | ---------- | ------ | ---- | ----- |
| `ES256K`        | ✅          | ✅      | ❌    | ❌     |
| `EdDSA:Ed25519` | ✅          | ✅      | ❌    | ❌     |


> [!IMPORTANT]
> In-memory signing using Secp256k1 **MUST** produce k-deterministic low-s signatures. Verification **must not require** low-s signatures 


## Key Management
Each SDK will implement a consistent and extensible _public interface_ for key management minimally providing concrete implementations of:

| Feature               | Typescript | Kotlin | Rust | Swift |
| --------------------- | ---------- | ------ | ---- | ----- |
| Key Manager Interface | ❌          | ✅      | ❌    | ❌     |
| In-Memory Key Manager | ❌          | ✅      | ❌    | ❌     |
| AWS KMS               | ❌          | ✅      | N/A  | N/A   |
| Device Enclave        | N/A        | ❌    | N/A  | ❌     |
| Keychain              | N/A        | ❌    | N/A  | ❌     |

Further, the key manager interface **must** be passed as an argument to _all_ public API methods that require key material. e.g.
* DID Creation
* Data Signing

> [!IMPORTANT]
> AWS KMS does **not** support `Ed25519`. At some point in the future, we can consider introducing support for key wrapping

> [!IMPORTANT]
> Consumers of our SDKs should be able to provide their own `KeyManager` implementations if desired

> [!NOTE]
> ⚠️ = implemented but no test vectors present

## `did:web`
| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| `Resolution` | ❌          | ❌      | ❌    | ❌     |

## `did:jwk`
| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| `Creation`   | ❌          | ❌      | ❌    | ❌     |
| `Resolution` | ❌          | ❌      | ❌    | ❌     |


## `did:dht`
| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| `Creation`   | ⚠️          | ⚠️      | ❌    | ❌     |
| `Resolution` | ⚠️          | ⚠️      | ❌    | ❌     |


## `did:key`
| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| `Creation`   | ⚠️          | ⚠️      | ❌    | ❌     |
| `Resolution` | ⚠️          | ⚠️      | ❌    | ❌     |

> [!IMPORTANT]
> `did:key` is included because it has been implemented in both Kotlin and Typescript. I'll be creating a Github issue soon to discuss when we think it makes sense to remove ION support from both SDKs


## `did:ion`
| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| `Creation`   | ⚠️          | ⚠️      | ❌    | ❌     |
| `Resolution` | ⚠️          | ⚠️      | ❌    | ❌     |

> [!IMPORTANT]
> `did:ion` is included because it has been implemented in both Kotlin and Typescript. I'll be creating a Github issue soon to discuss when we think it makes sense to remove ION support from both SDKs


## DID Document & Resolution Validation
| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| JSON Schema  | ❌          | ❌      | ❌    | ❌     |
| Common Error | ❌          | ❌      | ❌    | ❌     |


## W3C Verifiable Credential Data Model 1.1
| Feature             | Typescript | Kotlin | Rust | Swift |
| ------------------- | ---------- | ------ | ---- | ----- |
| Creation            | ❌          | ❌      | ❌    | ❌     |
| Signing as `vc-jwt` | ❌          | ❌      | ❌    | ❌     |
| Verification        | ❌          | ❌      | ❌    | ❌     |
| Validation          | ❌          | ❌      | ❌    | ❌     |

## W3C Verifiable Credential Data Model 2.0
| Feature             | Typescript | Kotlin | Rust | Swift |
| ------------------- | ---------- | ------ | ---- | ----- |
| Creation            | ❌          | ❌      | ❌    | ❌     |
| Signing as `vc-jwt` | ❌          | ❌      | ❌    | ❌     |
| Verification        | ❌          | ❌      | ❌    | ❌     |
| Validation          | ❌          | ❌      | ❌    | ❌     |

## SD-JWT / SD-JWT-VC
| Feature             | Typescript | Kotlin | Rust | Swift |
| ------------------- | ---------- | ------ | ---- | ----- |
| Creation            | ❌          | ❌      | ❌    | ❌     |
| Signing as `vc-jwt` | ❌          | ❌      | ❌    | ❌     |
| Verification        | ❌          | ❌      | ❌    | ❌     |
| Validation          | ❌          | ❌      | ❌    | ❌     |

## Presentation Definition V2
| Feature               | Typescript | Kotlin | Rust | Swift |
| --------------------- | ---------- | ------ | ---- | ----- |
| Concrete Type         | ❌          | ❌      | ❌    | ❌     |
| Validation            | ❌          | ❌      | ❌    | ❌     |
| Credential Evaluation | ❌          | ❌      | ❌    | ❌     |

## tbDEX Message

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| Validation   | ✅          | ✅      | ❌    | ❌     |
| Signing      | ✅          | ✅      | ❌    | ❌     |
| Verification | ✅          | ✅      | ❌    | ❌     |
| Parsing      | ✅          | ✅      | ❌    | ❌     |


## tbDEX Resource

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| Validation   | ✅          | ✅      | ❌    | ❌     |
| Signing      | ✅          | ✅      | ❌    | ❌     |
| Verification | ✅          | ✅      | ❌    | ❌     |
| Parsing      | ✅          | ✅      | ❌    | ❌     |


## tbDEX Offering Resource

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| Creation     | ✅          | ✅      | ❌    | ❌     |
| Validation   | ✅          | ✅      | ❌    | ❌     |
| Signing      | ✅          | ✅      | ❌    | ❌     |
| Verification | ✅          | ✅      | ❌    | ❌     |
| Parsing      | ✅          | ✅      | ❌    | ❌     |

## tbDEX RFQ Message

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| Creation     | ✅          | ✅      | ❌    | ❌     |
| Validation   | ✅          | ✅      | ❌    | ❌     |
| Signing      | ✅          | ✅      | ❌    | ❌     |
| Verification | ✅          | ✅      | ❌    | ❌     |
| Parsing      | ✅          | ✅      | ❌    | ❌     |

## tbDEX Quote Message

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| Creation     | ✅          | ✅      | ❌    | ❌     |
| Validation   | ✅          | ✅      | ❌    | ❌     |
| Signing      | ✅          | ✅      | ❌    | ❌     |
| Verification | ✅          | ✅      | ❌    | ❌     |
| Parsing      | ✅          | ✅      | ❌    | ❌     |

## tbDEX Order Message

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| Creation     | ✅          | ✅      | ❌    | ❌     |
| Validation   | ✅          | ✅      | ❌    | ❌     |
| Signing      | ✅          | ✅      | ❌    | ❌     |
| Verification | ✅          | ✅      | ❌    | ❌     |
| Parsing      | ✅          | ✅      | ❌    | ❌     |

## tbDEX Order-Status Message

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| Creation     | ✅          | ✅      | ❌    | ❌     |
| Validation   | ✅          | ✅      | ❌    | ❌     |
| Signing      | ✅          | ✅      | ❌    | ❌     |
| Verification | ✅          | ✅      | ❌    | ❌     |
| Parsing      | ✅          | ✅      | ❌    | ❌     |

## tbDEX Close Message

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| Creation     | ✅          | ✅      | ❌    | ❌     |
| Validation   | ✅          | ✅      | ❌    | ❌     |
| Signing      | ✅          | ✅      | ❌    | ❌     |
| Verification | ✅          | ✅      | ❌    | ❌     |
| Parsing      | ✅          | ✅      | ❌    | ❌     |
