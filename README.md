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
- [Test Vectors](#test-vectors)
  - [Usage](#usage)
    - [Local Dev](#local-dev)
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

This repo sets forth the development process and requirements for the following SDKs:

- [tbdex-js](https://github.com/TBD54566975/tbdex-js)
- [tbdex-kt](https://github.com/TBD54566975/tbdex-kt)
- [tbdex-rs](https://github.com/TBD54566975/tbdex-rs)
- [web5-js](https://github.com/TBD54566975/web5-js)
- [web5-kt](https://github.com/TBD54566975/web5-kt)
- [web5-rs](https://github.com/TBD54566975/web5-rs)

It contains the desired feature set for:

- [web5-js](https://github.com/TBD54566975/web5-js)
- [web5-kt](https://github.com/TBD54566975/web5-kt)
- [web5-rs](https://github.com/TBD54566975/web5-rs)

Github Issues and PRs created in this repo address topics that impact / span all SDKs we have in development. See section on [Feature Acceptance](#feature-acceptance) for information on submitting proposals for new features across all of our SDKs (e.g. implement Presentation Exchange v9)

## Known Unknowns

- We don't yet have a cemented path for supporting Android. It could either be through our pre-existing Kotlin SDK or bindings exposed via Rust

- We're not yet in a place where we can confidently state that a rust core can support a broad surface area of target languages. We're confident in pursuing rust development for the purposes of surfacing swift bindings.

## Requirements

### Feature Tracking

Feature tracking will take place entirely on GitHub through Github Issues and the [SDK development Github project](https://github.com/orgs/TBD54566975/projects/29).

---

Work items that impact all SDKs should be created as issues in _this_ Github Repo. e.g. "produce test vectors for `did:web`

---

Work items that are specific to an individual SDK should be created as an issue in the respective sdk's Github Repo

> [!IMPORTANT]
> Each individual SDK's DRI will be responsible for creating all of the relevant github issues in their respective repo by using the tables below. After all issues are created for a given feature, create a PR for this repo to delete the relevant table.

---

All relevant Github Issues will be tracked in the [SDK Development](https://github.com/orgs/TBD54566975/projects/29) Project. Issues can be associated to the project through the sidebar on an individual issue page. If a new issue is created in any of the SDK repos with one of the below feature labels, it will _automatically_ be added to the project board via GH project workflows. _New_ features lables should be added to the [workflows](https://github.com/orgs/TBD54566975/projects/29/workflows).

---

Work item progress is tracked using the `Status` attribute on a GH issue. This should automatically be reflected on the kanban view on the GH project. PRs should be linked to their respective issue via the PR description.

---

#### Labels

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
| `vc-json-schema`     | `#C86F42` | VC JSON Schema                             |
| `tbdex-message`      | `#70DB93` | tbDEX Message                              |
| `tbdex-resource`     | `#5B2C6F` | tbDEX Resource                             |
| `tbdex-offering`     | `#E59866` | tbDEX Offering Resource                    |
| `tbdex-rfq`          | `#1F618D` | tbDEX RFQ Message                          |
| `tbdex-quote`        | `#186A3B` | tbDEX Quote Message                        |
| `tbdex-order`        | `#28B463` | tbDEX Order Message                        |
| `tbdex-orderstatus`  | `#D68910` | tbDEX Order-Status Message                 |
| `tbdex-close`        | `#34495E` | tbDEX Close Message                        |
| `tbdex-server`       | `#3498DB` | HTTP server for tbDEX PFIs                 |
| `tbdex-client`       | `#E74C3C` | HTTP client for tbDEX wallets              |

> [!NOTE]
> This list will change over time as features are added or removed

---

#### Milestones

Github Repo Milestones will be used to track work for specific codenamed projects (e.g. Eagle, Pigeon, Black Swan, etc.).

The following milestones should exist in all relevant repos

| Milestone |
| --------- |
| `Eagle`   |
| `ABC`     |

> [!NOTE]
> This list will change over time as projects are added

### Feature Acceptance

Proposing new features that impact all SDKs will occur by creating a Github Issue in this repo. The Issue should include motivation or rationale in addition to any relevant reading material. New features will be discussed and decided upon during weekly syncs

> [!IMPORTANT]
> Language agnostic test vectors **must** be produced _prior_ to commencing implementation beyond the first SDK

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
| OSS License Check             | ❌         | ❌     | ✅   | ❌    |
| Security Scanning             | ❌         | ❌     | ⛔️  | ❌    |
| Static Analysis Linting/Style | ❌         | ❌     | ✅   | ❌    |
| Running Unit Tests            | ✅         | ✅     | ✅   | ❌    |
| Publishing Tests Reports      | ❌         | ❌     | ❌   | ❌    |
| Code Coverage (CodeCov)       | ❌         | ❌     | ❌   | ❌    |
| Publishing Artifacts          | ❓         | ❌     | ❌   | ❌    |
| Release Template Checklist    | ❌         | ❌     | ❌   | ❌    |
| Automated GH Release Tag      | ❌         | ❌     | ❌   | ❌    |
| Publishing API Reference Docs | ❌         | ✅     | ❌   | ❌    |
| Publish Example Feature Usage | ❌         | ✅     | ❌   | ❌    |

> [!CAUTION]
> Security scanning via Snyk is currently not supported in Rust

- GitHub Actions should run in secured runners
  - A secure, authoritative build environment ensures software is compiled and packaged in a controlled, tamper-resistant setting.
  - This mitigates the risk of introducing vulnerabilities or malicious code during the build process, whether through external attacks or compromised internal components.
  - These runners are going to be TBD-owned and self hosted
- Ideally the above table should be represented by a "Software Catalog" with all of our SDK statuses in real time.
  - The dashboard would be consuming the data sources (GitHub, CodeCov, Snyk, Npm and other registries etc.)
  - Tools like Grafana, Backstage, or even Jenkins (weather flag) could aggregate them

### CI/CD Statuses Dashboard

| SDK                             | CI Status                                                 | License                                                                                                                                           | License & Security Scanning                                                                                                   | OSSF Score                                        | SAST/Lint    | Unit Tests                                                           | Acceptance Tests | Release                                                                                                                                               | API Reference Docs                                                                                                                                                                                                  |
| ------------------------------- | --------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------- | ------------ | -------------------------------------------------------------------- | ---------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ![js-icon] [@web5/common]       | [![@web5/common:ci-img]][@web5/common:ci-url]             | [![@web5/common:ghlicense-img]][@web5/common:ghlicense-url] [![@web5/common:pkglicense-img]][@web5/common:pkglicense-url]                         | [![@web5/common:scan-img]][@web5/common:scan-url] [![@web5/common:fossa-img]][@web5/common:fossa-url]                         | [![@web5/common:ossf-img]][@web5/common:ossf-url] | ![todo-sast] | ![todo-unit] [![@web5/common:codecov-img]][@web5/common:codecov-url] | ![todo-e2e]      | [![@web5/common:ghtag-img]][@web5/common:ghtag-url] [![@web5/common:pkg-img]][@web5/common:pkg-url] ![todo-spdx] ![todo-cydx]                         | ![todo-docsci] ![todo-docscd]                                                                                                                                                                                       |
| ![js-icon] [@web5/crypto]       | [![@web5/crypto:ci-img]][@web5/crypto:ci-url]             | [![@web5/crypto:ghlicense-img]][@web5/crypto:ghlicense-url] [![@web5/crypto:pkglicense-img]][@web5/crypto:pkglicense-url]                         | [![@web5/crypto:scan-img]][@web5/crypto:scan-url] [![@web5/crypto:fossa-img]][@web5/crypto:fossa-url]                         | [![@web5/crypto:ossf-img]][@web5/crypto:ossf-url] | ![todo-sast] | ![todo-unit] [![@web5/crypto:codecov-img]][@web5/crypto:codecov-url] | ![todo-e2e]      | [![@web5/crypto:ghtag-img]][@web5/crypto:ghtag-url] [![@web5/crypto:pkg-img]][@web5/crypto:pkg-url] ![todo-spdx] ![todo-cydx]                         | ![todo-docsci] ![todo-docscd]                                                                                                                                                                                       |
| ![js-icon] [@web5/dids]         | [![@web5/dids:ci-img]][@web5/dids:ci-url]                 | [![@web5/dids:ghlicense-img]][@web5/dids:ghlicense-url] [![@web5/dids:pkglicense-img]][@web5/dids:pkglicense-url]                                 | [![@web5/dids:scan-img]][@web5/dids:scan-url] [![@web5/dids:fossa-img]][@web5/dids:fossa-url]                                 | [![@web5/dids:ossf-img]][@web5/dids:ossf-url]     | ![todo-sast] | ![todo-unit] [![@web5/dids:codecov-img]][@web5/dids:codecov-url]     | ![todo-e2e]      | [![@web5/dids:ghtag-img]][@web5/dids:ghtag-url] [![@web5/dids:pkg-img]][@web5/dids:pkg-url] ![todo-spdx] ![todo-cydx]                                 | ![todo-docsci] ![todo-docscd]                                                                                                                                                                                       |
| ![js-icon] [@tbdex/http-client] | [![@tbdex/http-client:ci-img]][@tbdex/http-client:ci-url] | [![@tbdex/http-client:ghlicense-img]][@tbdex/http-client:ghlicense-url] [![@tbdex/http-client:pkglicense-img]][@tbdex/http-client:pkglicense-url] | [![@tbdex/http-client:scan-img]][@tbdex/http-client:scan-url] [![@tbdex/http-client:fossa-img]][@tbdex/http-client:fossa-url] | ![todo-ossf]                                      | ![todo-sast] | ![todo-unit] ![todo-codecov]                                         | ![todo-e2e]      | [![@tbdex/http-client:ghtag-img]][@tbdex/http-client:ghtag-url] [![@tbdex/http-client:pkg-img]][@tbdex/http-client:pkg-url] ![todo-spdx] ![todo-cydx] | [![@tbdex/http-client:apidocsci-img]][@tbdex/http-client:apidocsci-url] [![@tbdex/http-client:apidocscd-img]][@tbdex/http-client:apidocscd-url] [![@tbdex/http-client:apidocs-img]][@tbdex/http-client:apidocs-url] |
| ![js-icon] [@tbdex/http-server] | [![@tbdex/http-server:ci-img]][@tbdex/http-server:ci-url] | [![@tbdex/http-server:ghlicense-img]][@tbdex/http-server:ghlicense-url] [![@tbdex/http-server:pkglicense-img]][@tbdex/http-server:pkglicense-url] | [![@tbdex/http-server:scan-img]][@tbdex/http-server:scan-url] [![@tbdex/http-server:fossa-img]][@tbdex/http-server:fossa-url] | ![todo-ossf]                                      | ![todo-sast] | ![todo-unit] ![todo-codecov]                                         | ![todo-e2e]      | [![@tbdex/http-server:ghtag-img]][@tbdex/http-server:ghtag-url] [![@tbdex/http-server:pkg-img]][@tbdex/http-server:pkg-url] ![todo-spdx] ![todo-cydx] | [![@tbdex/http-server:apidocsci-img]][@tbdex/http-server:apidocsci-url] [![@tbdex/http-server:apidocscd-img]][@tbdex/http-server:apidocscd-url] [![@tbdex/http-server:apidocs-img]][@tbdex/http-server:apidocs-url] |
| ![js-icon] [@tbdex/protocol]    | [![@tbdex/protocol:ci-img]][@tbdex/protocol:ci-url]       | [![@tbdex/protocol:ghlicense-img]][@tbdex/protocol:ghlicense-url] [![@tbdex/protocol:pkglicense-img]][@tbdex/protocol:pkglicense-url]             | [![@tbdex/protocol:scan-img]][@tbdex/protocol:scan-url] [![@tbdex/protocol:fossa-img]][@tbdex/protocol:fossa-url]             | ![todo-ossf]                                      | ![todo-sast] | ![todo-unit] ![todo-codecov]                                         | ![todo-e2e]      | [![@tbdex/protocol:ghtag-img]][@tbdex/protocol:ghtag-url] [![@tbdex/protocol:pkg-img]][@tbdex/protocol:pkg-url] ![todo-spdx] ![todo-cydx]             | [![@tbdex/protocol:apidocsci-img]][@tbdex/protocol:apidocsci-url] [![@tbdex/protocol:apidocscd-img]][@tbdex/protocol:apidocscd-url] [![@tbdex/protocol:apidocs-img]][@tbdex/protocol:apidocs-url]                   |
| ![kt-icon] [tbdex-kt]           | [![tbdex-kt:ci-img]][tbdex-kt:ci-url]                     | [![tbdex-kt:ghlicense-img]][tbdex-kt:ghlicense-url] ![tbdex-kt:pkglicense-img]                                                                    | [![tbdex-kt:scan-img]][tbdex-kt:scan-url] [![tbdex-kt:fossa-img]][tbdex-kt:fossa-url]                                         | ![todo-ossf]                                      | ![todo-sast] | ![todo-unit] [![tbdex-kt:codecov-img]][tbdex-kt:codecov-url]         | ![todo-e2e]      | [![tbdex-kt:ghtag-img]][tbdex-kt:ghtag-url] [![tbdex-kt:pkg-img]][tbdex-kt:pkg-url] ![todo-spdx] ![todo-cydx]                                         | [![tbdex-kt:apidocsci-img]][tbdex-kt:apidocsci-url] [![tbdex-kt:apidocscd-img]][tbdex-kt:apidocscd-url] [![tbdex-kt:apidocs-img]][tbdex-kt:apidocs-url]                                                             |

<!-- MISC BADGES -->

[js-icon]: https://img.shields.io/badge/-js-F7DF1E?style=flat-square
[kt-icon]: https://img.shields.io/badge/-kt-7F52FF?style=flat-square
[todo-ossf]: https://img.shields.io/badge/ossf-todo-indigo?style=flat-square
[todo-unit]: https://img.shields.io/badge/unit-todo-indigo?style=flat-square
[todo-codecov]: https://img.shields.io/badge/codecov-todo-indigo?style=flat-square
[todo-sast]: https://img.shields.io/badge/sast-todo-indigo?style=flat-square
[todo-lint]: https://img.shields.io/badge/lint-todo-indigo?style=flat-square
[todo-docsci]: https://img.shields.io/badge/docs%20ci-todo-indigo?style=flat-square
[todo-docscd]: https://img.shields.io/badge/docs%20publish-todo-indigo?style=flat-square
[todo-spdx]: https://img.shields.io/badge/spdx-todo-indigo?style=flat-square
[todo-cydx]: https://img.shields.io/badge/cydx-todo-indigo?style=flat-square
[todo-e2e]: https://img.shields.io/badge/e2e-todo-indigo?style=flat-square

<!-- @web5/common -->

[@web5/common]: https://github.com/TBD54566975/web5-js/tree/main/packages/common
[@web5/common:ci-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/web5-js/tests-ci.yml?branch=main&logo=github&label=ci&logoColor=FFFFFF&style=flat-square
[@web5/common:ci-url]: https://github.com/TBD54566975/web5-js/actions/workflows/tests-ci.yml
[@web5/common:ghlicense-img]: https://img.shields.io/github/license/TBD54566975/web5-js?style=flat-square&logo=github&color=4c1&label=gh
[@web5/common:ghlicense-url]: https://github.com/TBD54566975/web5-js/blob/main/LICENSE
[@web5/common:scan-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/web5-js/security.yml?branch=main&logo=github&label=scan&logoColor=FFFFFF&style=flat-square
[@web5/common:scan-url]: https://github.com/TBD54566975/web5-js/actions/workflows/security.yml
[@web5/common:ossf-img]: https://img.shields.io/ossf-scorecard/github.com/TBD54566975/web5-js?label=ossf&style=flat-square
[@web5/common:ossf-url]: https://securityscorecards.dev/viewer/?uri=github.com/TBD54566975/web5-js
[@web5/common:codecov-img]: https://img.shields.io/codecov/c/gh/TBD54566975/web5-js/main?label=codecov&style=flat-square&token=YI87CKF1LI
[@web5/common:codecov-url]: https://codecov.io/github/TBD54566975/web5-js
[@web5/common:ghtag-img]: https://img.shields.io/github/v/release/TBD54566975/web5-js?logo=github&label=tag&style=flat-square&color=4c1
[@web5/common:ghtag-url]: https://github.com/TBD54566975/web5-js/releases
[@web5/common:pkg-img]: https://img.shields.io/npm/v/@web5/common.svg?style=flat-square&logo=npm&logoColor=FFFFFF&color=F7DF1E&santize=true
[@web5/common:pkg-url]: https://www.npmjs.com/package/@web5/common
[@web5/common:pkglicense-img]: https://img.shields.io/npm/l/@web5/common.svg?style=flat-square&color=F7DF1E&logo=npm&logoColor=FFFFFF&santize=true&label=npm
[@web5/common:pkglicense-url]: https://www.npmjs.com/package/@web5/common
[@web5/common:fossa-img]: https://app.fossa.com/api/projects/custom%2B588%2Fgithub.com%2FTBD54566975%2Fweb5-js.svg?type=small
[@web5/common:fossa-url]: https://app.fossa.com/projects/custom%2B588%2Fgithub.com%2FTBD54566975%2Fweb5-js?ref=badge_small

<!-- @web5/crypto -->

[@web5/crypto]: https://github.com/TBD54566975/web5-js/tree/main/packages/crypto
[@web5/crypto:ci-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/web5-js/tests-ci.yml?branch=main&logo=github&label=ci&logoColor=FFFFFF&style=flat-square
[@web5/crypto:ci-url]: https://github.com/TBD54566975/web5-js/actions/workflows/tests-ci.yml
[@web5/crypto:ghlicense-img]: https://img.shields.io/github/license/TBD54566975/web5-js?style=flat-square&logo=github&color=4c1&label=gh
[@web5/crypto:ghlicense-url]: https://github.com/TBD54566975/web5-js/blob/main/LICENSE
[@web5/crypto:scan-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/web5-js/security.yml?branch=main&logo=github&label=scan&logoColor=FFFFFF&style=flat-square
[@web5/crypto:scan-url]: https://github.com/TBD54566975/web5-js/actions/workflows/security.yml
[@web5/crypto:ossf-img]: https://img.shields.io/ossf-scorecard/github.com/TBD54566975/web5-js?label=ossf&style=flat-square
[@web5/crypto:ossf-url]: https://securityscorecards.dev/viewer/?uri=github.com/TBD54566975/web5-js
[@web5/crypto:codecov-img]: https://img.shields.io/codecov/c/gh/TBD54566975/web5-js/main?label=codecov&style=flat-square&token=YI87CKF1LI
[@web5/crypto:codecov-url]: https://codecov.io/github/TBD54566975/web5-js
[@web5/crypto:ghtag-img]: https://img.shields.io/github/v/release/TBD54566975/web5-js?logo=github&label=tag&style=flat-square&color=4c1
[@web5/crypto:ghtag-url]: https://github.com/TBD54566975/web5-js/releases
[@web5/crypto:pkg-img]: https://img.shields.io/npm/v/@web5/crypto.svg?style=flat-square&logo=npm&logoColor=FFFFFF&color=F7DF1E&santize=true
[@web5/crypto:pkg-url]: https://www.npmjs.com/package/@web5/crypto
[@web5/crypto:pkglicense-img]: https://img.shields.io/npm/l/@web5/crypto.svg?style=flat-square&color=F7DF1E&logo=npm&logoColor=FFFFFF&santize=true&label=npm
[@web5/crypto:pkglicense-url]: https://www.npmjs.com/package/@web5/crypto
[@web5/crypto:fossa-img]: https://app.fossa.com/api/projects/custom%2B588%2Fgithub.com%2FTBD54566975%2Fweb5-js.svg?type=small
[@web5/crypto:fossa-url]: https://app.fossa.com/projects/custom%2B588%2Fgithub.com%2FTBD54566975%2Fweb5-js?ref=badge_small

<!-- @web5/dids -->

[@web5/dids]: https://github.com/TBD54566975/web5-js/tree/main/packages/dids
[@web5/dids:ci-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/web5-js/tests-ci.yml?branch=main&logo=github&label=ci&logoColor=FFFFFF&style=flat-square
[@web5/dids:ci-url]: https://github.com/TBD54566975/web5-js/actions/workflows/tests-ci.yml
[@web5/dids:ghlicense-img]: https://img.shields.io/github/license/TBD54566975/web5-js?style=flat-square&logo=github&color=4c1&label=gh
[@web5/dids:ghlicense-url]: https://github.com/TBD54566975/web5-js/blob/main/LICENSE
[@web5/dids:scan-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/web5-js/security.yml?branch=main&logo=github&label=scan&logoColor=FFFFFF&style=flat-square
[@web5/dids:scan-url]: https://github.com/TBD54566975/web5-js/actions/workflows/security.yml
[@web5/dids:ossf-img]: https://img.shields.io/ossf-scorecard/github.com/TBD54566975/web5-js?label=ossf&style=flat-square
[@web5/dids:ossf-url]: https://securityscorecards.dev/viewer/?uri=github.com/TBD54566975/web5-js
[@web5/dids:codecov-img]: https://img.shields.io/codecov/c/gh/TBD54566975/web5-js/main?label=codecov&style=flat-square&token=YI87CKF1LI
[@web5/dids:codecov-url]: https://codecov.io/github/TBD54566975/web5-js
[@web5/dids:ghtag-img]: https://img.shields.io/github/v/release/TBD54566975/web5-js?logo=github&label=tag&style=flat-square&color=4c1
[@web5/dids:ghtag-url]: https://github.com/TBD54566975/web5-js/releases
[@web5/dids:pkg-img]: https://img.shields.io/npm/v/@web5/dids.svg?style=flat-square&logo=npm&logoColor=FFFFFF&color=F7DF1E&santize=true
[@web5/dids:pkg-url]: https://www.npmjs.com/package/@web5/dids
[@web5/dids:pkglicense-img]: https://img.shields.io/npm/l/@web5/dids.svg?style=flat-square&color=F7DF1E&logo=npm&logoColor=FFFFFF&santize=true&label=npm
[@web5/dids:pkglicense-url]: https://www.npmjs.com/package/@web5/dids
[@web5/dids:fossa-img]: https://app.fossa.com/api/projects/custom%2B588%2Fgithub.com%2FTBD54566975%2Fweb5-js.svg?type=small
[@web5/dids:fossa-url]: https://app.fossa.com/projects/custom%2B588%2Fgithub.com%2FTBD54566975%2Fweb5-js?ref=badge_small

<!-- @tbdex/protocol -->

[@tbdex/protocol]: https://github.com/TBD54566975/tbdex-js/tree/main/packages/protocol
[@tbdex/protocol:ci-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/tbdex-js/integrity-check.yml?branch=main&logo=github&label=ci&logoColor=FFFFFF&style=flat-square
[@tbdex/protocol:ci-url]: https://github.com/TBD54566975/tbdex-js/actions/workflows/integrity-check.yml
[@tbdex/protocol:ghlicense-img]: https://img.shields.io/github/license/TBD54566975/tbdex-js?style=flat-square&logo=github&color=4c1&label=gh
[@tbdex/protocol:ghlicense-url]: https://github.com/TBD54566975/tbdex-js/blob/main/LICENSE
[@tbdex/protocol:scan-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/tbdex-js/security.yml?branch=main&logo=github&label=scan&logoColor=FFFFFF&style=flat-square
[@tbdex/protocol:scan-url]: https://github.com/TBD54566975/tbdex-js/actions/workflows/security.yml
[@tbdex/protocol:ghtag-img]: https://img.shields.io/github/v/release/TBD54566975/tbdex-js?logo=github&label=tag&style=flat-square&color=4c1
[@tbdex/protocol:ghtag-url]: https://github.com/TBD54566975/tbdex-js/releases
[@tbdex/protocol:pkg-img]: https://img.shields.io/npm/v/@tbdex/protocol.svg?style=flat-square&logo=npm&logoColor=FFFFFF&color=F7DF1E&santize=true
[@tbdex/protocol:pkg-url]: https://www.npmjs.com/package/@tbdex/protocol
[@tbdex/protocol:pkglicense-img]: https://img.shields.io/npm/l/@tbdex/protocol.svg?style=flat-square&color=F7DF1E&logo=npm&logoColor=FFFFFF&santize=true&label=npm
[@tbdex/protocol:pkglicense-url]: https://www.npmjs.com/package/@tbdex/protocol
[@tbdex/protocol:apidocsci-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/tbdex-js/docs-ci.yml?branch=main&logo=github&label=docs%20ci&logoColor=FFFFFF&style=flat-square
[@tbdex/protocol:apidocsci-url]: https://github.com/TBD54566975/tbdex-js/actions/workflows/docs-ci.yml
[@tbdex/protocol:apidocscd-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/tbdex-js/docs-publish.yml?branch=main&logo=github&label=docs%20publish&logoColor=FFFFFF&style=flat-square
[@tbdex/protocol:apidocscd-url]: https://github.com/TBD54566975/tbdex-js/actions/workflows/docs-publish.yml
[@tbdex/protocol:apidocs-img]: https://img.shields.io/badge/reference_docs-7F52FF?style=flat-square
[@tbdex/protocol:apidocs-url]: https://tbd54566975.github.io/tbdex-js/modules/_tbdex_protocol.html
[@tbdex/protocol:fossa-img]: https://app.fossa.com/api/projects/custom%2B588%2Fgithub.com%2FTBD54566975%2Ftbdex-js.svg?type=small
[@tbdex/protocol:fossa-url]: https://app.fossa.com/projects/custom%2B588%2Fgithub.com%2FTBD54566975%2Ftbdex-js?ref=badge_small

<!-- @tbdex/http-client -->

[@tbdex/http-client]: https://github.com/TBD54566975/tbdex-js/tree/main/packages/http-client
[@tbdex/http-client:ci-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/tbdex-js/integrity-check.yml?branch=main&logo=github&label=ci&logoColor=FFFFFF&style=flat-square
[@tbdex/http-client:ci-url]: https://github.com/TBD54566975/tbdex-js/actions/workflows/integrity-check.yml
[@tbdex/http-client:ghlicense-img]: https://img.shields.io/github/license/TBD54566975/tbdex-js?style=flat-square&logo=github&color=4c1&label=gh
[@tbdex/http-client:ghlicense-url]: https://github.com/TBD54566975/tbdex-js/blob/main/LICENSE
[@tbdex/http-client:scan-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/tbdex-js/security.yml?branch=main&logo=github&label=scan&logoColor=FFFFFF&style=flat-square
[@tbdex/http-client:scan-url]: https://github.com/TBD54566975/tbdex-js/actions/workflows/security.yml
[@tbdex/http-client:ghtag-img]: https://img.shields.io/github/v/release/TBD54566975/tbdex-js?logo=github&label=tag&style=flat-square&color=4c1
[@tbdex/http-client:ghtag-url]: https://github.com/TBD54566975/tbdex-js/releases
[@tbdex/http-client:pkg-img]: https://img.shields.io/npm/v/@tbdex/http-client.svg?style=flat-square&logo=npm&logoColor=FFFFFF&color=F7DF1E&santize=true
[@tbdex/http-client:pkg-url]: https://www.npmjs.com/package/@tbdex/http-client
[@tbdex/http-client:pkglicense-img]: https://img.shields.io/npm/l/@tbdex/http-client.svg?style=flat-square&color=F7DF1E&logo=npm&logoColor=FFFFFF&santize=true&label=npm
[@tbdex/http-client:pkglicense-url]: https://www.npmjs.com/package/@tbdex/http-client
[@tbdex/http-client:apidocsci-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/tbdex-js/docs-ci.yml?branch=main&logo=github&label=docs%20ci&logoColor=FFFFFF&style=flat-square
[@tbdex/http-client:apidocsci-url]: https://github.com/TBD54566975/tbdex-js/actions/workflows/docs-ci.yml
[@tbdex/http-client:apidocscd-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/tbdex-js/docs-publish.yml?branch=main&logo=github&label=docs%20publish&logoColor=FFFFFF&style=flat-square
[@tbdex/http-client:apidocscd-url]: https://github.com/TBD54566975/tbdex-js/actions/workflows/docs-publish.yml
[@tbdex/http-client:apidocs-img]: https://img.shields.io/badge/reference_docs-7F52FF?style=flat-square
[@tbdex/http-client:apidocs-url]: https://tbd54566975.github.io/tbdex-js/modules/_tbdex_http_client.html
[@tbdex/http-client:fossa-img]: https://app.fossa.com/api/projects/custom%2B588%2Fgithub.com%2FTBD54566975%2Ftbdex-js.svg?type=small
[@tbdex/http-client:fossa-url]: https://app.fossa.com/projects/custom%2B588%2Fgithub.com%2FTBD54566975%2Ftbdex-js?ref=badge_small

<!-- @tbdex/http-server -->

[@tbdex/http-server]: https://github.com/TBD54566975/tbdex-js/tree/main/packages/http-server
[@tbdex/http-server:ci-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/tbdex-js/integrity-check.yml?branch=main&logo=github&label=ci&logoColor=FFFFFF&style=flat-square
[@tbdex/http-server:ci-url]: https://github.com/TBD54566975/tbdex-js/actions/workflows/integrity-check.yml
[@tbdex/http-server:ghlicense-img]: https://img.shields.io/github/license/TBD54566975/tbdex-js?style=flat-square&logo=github&color=4c1&label=gh
[@tbdex/http-server:ghlicense-url]: https://github.com/TBD54566975/tbdex-js/blob/main/LICENSE
[@tbdex/http-server:scan-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/tbdex-js/security.yml?branch=main&logo=github&label=scan&logoColor=FFFFFF&style=flat-square
[@tbdex/http-server:scan-url]: https://github.com/TBD54566975/tbdex-js/actions/workflows/security.yml
[@tbdex/http-server:ghtag-img]: https://img.shields.io/github/v/release/TBD54566975/tbdex-js?logo=github&label=tag&style=flat-square&color=4c1
[@tbdex/http-server:ghtag-url]: https://github.com/TBD54566975/tbdex-js/releases
[@tbdex/http-server:pkg-img]: https://img.shields.io/npm/v/@tbdex/http-server.svg?style=flat-square&logo=npm&logoColor=FFFFFF&color=F7DF1E&santize=true
[@tbdex/http-server:pkg-url]: https://www.npmjs.com/package/@tbdex/http-server
[@tbdex/http-server:pkglicense-img]: https://img.shields.io/npm/l/@tbdex/http-server.svg?style=flat-square&color=F7DF1E&logo=npm&logoColor=FFFFFF&santize=true&label=npm
[@tbdex/http-server:pkglicense-url]: https://www.npmjs.com/package/@tbdex/http-server
[@tbdex/http-server:apidocsci-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/tbdex-js/docs-ci.yml?branch=main&logo=github&label=docs%20ci&logoColor=FFFFFF&style=flat-square
[@tbdex/http-server:apidocsci-url]: https://github.com/TBD54566975/tbdex-js/actions/workflows/docs-ci.yml
[@tbdex/http-server:apidocscd-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/tbdex-js/docs-publish.yml?branch=main&logo=github&label=docs%20publish&logoColor=FFFFFF&style=flat-square
[@tbdex/http-server:apidocscd-url]: https://github.com/TBD54566975/tbdex-js/actions/workflows/docs-publish.yml
[@tbdex/http-server:apidocs-img]: https://img.shields.io/badge/reference_docs-7F52FF?style=flat-square
[@tbdex/http-server:apidocs-url]: https://tbd54566975.github.io/tbdex-js/modules/_tbdex_http_server.html
[@tbdex/http-server:fossa-img]: https://app.fossa.com/api/projects/custom%2B588%2Fgithub.com%2FTBD54566975%2Ftbdex-js.svg?type=small
[@tbdex/http-server:fossa-url]: https://app.fossa.com/projects/custom%2B588%2Fgithub.com%2FTBD54566975%2Ftbdex-js?ref=badge_small

<!-- tbdex-kt -->

[tbdex-kt]: https://github.com/TBD54566975/tbdex-kt/tree/main/httpclient
[tbdex-kt:ci-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/tbdex-kt/ci.yaml?branch=main&logo=github&label=ci&logoColor=FFFFFF&style=flat-square
[tbdex-kt:ci-url]: https://github.com/TBD54566975/tbdex-kt/actions/workflows/ci.yaml
[tbdex-kt:ghlicense-img]: https://img.shields.io/github/license/TBD54566975/tbdex-kt?style=flat-square&logo=github&color=4c1&label=gh
[tbdex-kt:ghlicense-url]: https://github.com/TBD54566975/tbdex-kt/blob/main/LICENSE
[tbdex-kt:scan-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/tbdex-kt/security.yaml?branch=main&logo=github&label=scan&logoColor=FFFFFF&style=flat-square
[tbdex-kt:scan-url]: https://github.com/TBD54566975/tbdex-kt/actions/workflows/security.yaml
[tbdex-kt:codecov-img]: https://img.shields.io/codecov/c/gh/TBD54566975/tbdex-kt/main?label=codecov&style=flat-square&token=YI87CKF1LI
[tbdex-kt:codecov-url]: https://codecov.io/github/TBD54566975/tbdex-kt
[tbdex-kt:ghtag-img]: https://img.shields.io/github/v/release/TBD54566975/tbdex-kt?logo=github&label=tag&style=flat-square&color=4c1
[tbdex-kt:ghtag-url]: https://github.com/TBD54566975/tbdex-kt/releases
[tbdex-kt:pkg-img]: https://img.shields.io/jitpack/version/com.github.TBD54566975/tbdex-kt?style=flat-square&logo=jitpack
[tbdex-kt:pkg-url]: https://jitpack.io/#TBD54566975/tbdex-kt
[tbdex-kt:pkglicense-img]: https://img.shields.io/badge/mvn-todo-indigo?style=flat-square&logo=apachemaven&logoColor=FFFFFF&santize=true
[tbdex-kt:apidocsci-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/tbdex-kt/ci.yaml?branch=main&logo=github&label=docs%20ci&logoColor=FFFFFF&style=flat-square
[tbdex-kt:apidocsci-url]: https://github.com/TBD54566975/tbdex-kt/actions/workflows/ci.yaml
[tbdex-kt:apidocscd-img]: https://img.shields.io/github/actions/workflow/status/TBD54566975/tbdex-kt/docs.yaml?branch=main&logo=github&label=docs%20publish&logoColor=FFFFFF&style=flat-square
[tbdex-kt:apidocscd-url]: https://github.com/TBD54566975/tbdex-kt/actions/workflows/docs.yaml
[tbdex-kt:apidocs-img]: https://img.shields.io/badge/reference_docs-7F52FF?style=flat-square
[tbdex-kt:apidocs-url]: https://tbd54566975.github.io/tbdex-kt/index.html
[tbdex-kt:fossa-img]: https://app.fossa.com/api/projects/custom%2B588%2Fgithub.com%2FTBD54566975%2Ftbdex-kt.svg?type=small
[tbdex-kt:fossa-url]: https://app.fossa.com/projects/custom%2B588%2Fgithub.com%2FTBD54566975%2Ftbdex-kt?ref=badge_small

### Publishing Artifacts

Each SDK will be published to the most widely adopted registry/repository for the respective language

| SDK        | Repository          |
| ---------- | ------------------- |
| Typescript | npm                 |
| Kotlin     | maven central       |
| Rust       | crates.io           |
| Swift      | swift package index |

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

> [!IMPORTANT]
> Producing API reference documentation is the responsibility of an _implementer_

### Example Feature Usage

Each SDK will **publish** example usage for _each_ implemented feature. This can either be included as a part of API reference documentation _or_ published separately

## Test Vectors

Test vectors ensure interoporability of features across SDKs and language implementations by providing common test cases with an input and expected output pair. They include both success and failure cases that can be vectorized.

This repo serves as the home for all web5 feature related vectors. They are available in the [web5-test-vectors](./web5-test-vectors/) directory and hosted on [Github Pages](https://tbd54566975.github.io/sdk-development/web5-test-vectors).

The `tbdex` repo houses tbdex feature related vectors. They are available in the [test-vectors](https://github.com/TBD54566975/tbdex/test-vectors) directory and hosted on [Github Pages](https://tbdex.dev/).

### Usage

#### Local Dev

SDK implementers should import vectors in order to test their implementation. The recommended pathway to consume them is as follows:

Fetch the vector and read it into a data model representing the vector structure or a JSON object like so:

```kt
// for web5 vectors
val stream = URL("https://tbd54566975.github.io/sdk-development/web5-test-vectors/did-jwk/resolve.json").openStream()
val vectorsJson = BufferedReader(InputStreamReader(stream)).readText()
return Json.jsonMapper.readTree(vectorsJson)

// for tbdex vectors
val stream = URL("https://tbdex.dev/test-vectors/resources/marshal.json").openStream()
val vectorsJson = BufferedReader(InputStreamReader(stream)).readText()
return Json.jsonMapper.readTree(vectorsJson)
```

The data model or JSON object can then be used in the implementer's unit testing framework of choice.

#### Adding/Updating Vectors

New test vectors should follow the standard [vector structure](./web5-test-vectors/vectors.schema.json). Vectors are automatically validated against the JSON schema via CI.

Create a PR in this repo for web5 vectors, or in [`tbdex`](https://github.com/TBD54566975/tbdex) for tbdex vectors with the proposed changes or additions.

#### Feature Completeness By SDK

Test vectors are also used to determine feature completeness via our [test harness](./test-harness/README.md). Results of test harness runs can be found [here](https://tbd54566975.github.io/sdk-development/).

## Web5 SDK Features

### Cryptographic Digital Signature Algorithms (DSA)

| Algorithm       | Typescript | Kotlin | Rust | Swift |
| --------------- | ---------- | ------ | ---- | ----- |
| `ES256K`        | ✅         | ✅     | ✅   | ❌    |
| `EdDSA:Ed25519` | ✅         | ✅     | ✅   | ❌    |

> [!IMPORTANT]
> In-memory signing using Secp256k1 **MUST** produce k-deterministic low-s signatures. Verification **must not require** low-s signatures

### Key Management

Each SDK will implement a consistent and extensible _public interface_ for key management minimally providing concrete implementations of:

| Feature               | Typescript | Kotlin | Rust | Swift |
| --------------------- | ---------- | ------ | ---- | ----- |
| Key Manager Interface | ❌         | ✅     | ✅   | ❌    |
| In-Memory Key Manager | ❌         | ✅     | ✅   | ❌    |
| AWS KMS               | ❌         | ✅     | N/A  | N/A   |
| Device Enclave        | N/A        | ❌     | N/A  | ❌    |
| Keychain              | N/A        | ❌     | N/A  | ❌    |

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
| `Resolution` | ❌         | ❌     | ⚠️   | ❌    |

### [`did:jwk`](https://github.com/quartzjer/did-jwk/blob/main/spec.md)

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| `Creation`   | ❌         | ❌     | ⚠️   | ❌    |
| `Resolution` | ❌         | ❌     | ⚠️   | ❌    |

### [`did:dht`](https://tbd54566975.github.io/did-dht-method/)

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| `Creation`   | ⚠️         | ⚠️     | ❌   | ❌    |
| `Resolution` | ⚠️         | ⚠️     | ❌   | ❌    |

### [`did:key`](https://w3c-ccg.github.io/did-method-key/)

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| `Creation`   | ⚠️         | ⚠️     | ⚠️   | ❌    |
| `Resolution` | ⚠️         | ⚠️     | ⚠️   | ❌    |

> [!IMPORTANT] > `did:key` is included because it has been implemented in both Kotlin and Typescript. I'll be creating a Github issue soon to discuss when we think it makes sense to remove ION support from both SDKs

### [`did:ion`](https://identity.foundation/sidetree/spec)

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| `Creation`   | ⚠️         | ⚠️     | ❌   | ❌    |
| `Resolution` | ⚠️         | ⚠️     | ❌   | ❌    |

> [!IMPORTANT] > `did:ion` is included because it has been implemented in both Kotlin and Typescript. I'll be creating a Github issue soon to discuss when we think it makes sense to remove ION support from both SDKs

### [DID Document](https://www.w3.org/TR/did-core/) & [Resolution Validation](https://w3c-ccg.github.io/did-resolution/)

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| Common Error | ❌         | ❌     | ❌   | ❌    |

### [W3C Verifiable Credential Data Model 1.1](https://www.w3.org/TR/vc-data-model)

| Feature                              | Typescript | Kotlin | Rust | Swift |
| ------------------------------------ | ---------- | ------ | ---- | ----- |
| Creation                             | ✅         | ✅     | ❌   | ❌    |
| Signing as `vc-jwt`                  | ✅         | ✅     | ❌   | ❌    |
| Verification                         | ✅         | ✅     | ❌   | ❌    |
| Validation                           | ✅         | ✅     | ❌   | ❌    |
| Verifiable Presentations as `vp-jwt` | ⚠️         | ⚠️     | ❌   | ❌    |

### [W3C Verifiable Credential Data Model 2.0](https://www.w3.org/TR/vc-data-model-2.0/)

| Feature                   | Typescript | Kotlin | Rust | Swift |
| ------------------------- | ---------- | ------ | ---- | ----- |
| Creation                  | ❌         | ❌     | ❌   | ❌    |
| Signing as `vc-jose-cose` | ❌         | ❌     | ❌   | ❌    |
| Verification              | ❌         | ❌     | ❌   | ❌    |
| Validation                | ❌         | ❌     | ❌   | ❌    |

### [SD-JWT](https://datatracker.ietf.org/doc/draft-ietf-oauth-selective-disclosure-jwt/) / [SD-JWT-VC](https://datatracker.ietf.org/doc/draft-ietf-oauth-sd-jwt-vc/)

| Feature      | Typescript | Kotlin | Rust | Swift |
| ------------ | ---------- | ------ | ---- | ----- |
| Creation     | ❌         | ❌     | ❌   | ❌    |
| Signing      | ❌         | ❌     | ❌   | ❌    |
| Verification | ❌         | ❌     | ❌   | ❌    |
| Validation   | ❌         | ❌     | ❌   | ❌    |

### [Bitstring Status List](https://w3c.github.io/vc-bitstring-status-list/)

| Feature                   | Typescript | Kotlin | Rust | Swift |
| ------------------------- | ---------- | ------ | ---- | ----- |
| Creation using `vc-jwt`   | ❌         | ❌     | ❌   | ❌    |
| Validation using `vc-jwt` | ❌         | ❌     | ❌   | ❌    |

### [VC JSON Schema](https://www.w3.org/TR/vc-json-schema/)

| Feature                           | Typescript | Kotlin | Rust | Swift |
| --------------------------------- | ---------- | ------ | ---- | ----- |
| Creation `JsonSchema`             | ❌         | ❌     | ❌   | ❌    |
| Creation `JsonSchemaCredential`   | ❌         | ❌     | ❌   | ❌    |
| Validation `JsonSchema`           | ❌         | ❌     | ❌   | ❌    |
| Validation `JsonSchemaCredential` | ❌         | ❌     | ❌   | ❌    |

### [Presentation Exchange V2](https://identity.foundation/presentation-exchange/spec/v2.0.0/)

| Feature                                                                                                                  | Typescript | Kotlin | Rust | Swift |
| ------------------------------------------------------------------------------------------------------------------------ | ---------- | ------ | ---- | ----- |
| Concrete Type                                                                                                            | ✅         | ✅     | ❌   | ❌    |
| Validation                                                                                                               | ✅         | ✅     | ❌   | ❌    |
| Credential Evaluation                                                                                                    | ✅         | ✅     | ❌   | ❌    |
| [Predicates](https://identity.foundation/presentation-exchange/spec/v2.0.0/#predicate-feature)                           | ✅         | ✅     | ❌   | ❌    |
| [Relational Constraints](https://identity.foundation/presentation-exchange/spec/v2.0.0/#relational-constraint-feature)   | ✅         | ❌     | ❌   | ❌    |
| [Credential Status](https://identity.foundation/presentation-exchange/spec/v2.0.0/#credential-status-constraint-feature) | ❌         | ❌     | ❌   | ❌    |
