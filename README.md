# Web5 Spec

This repo sets forth the development process, requirements, and the desired feature for Web5 SDKs.

## Feature Tracking

Feature tracking will take place entirely on GitHub through Github Issues and the [SDK development Github project](https://github.com/orgs/TBD54566975/projects/29).

## Feature Acceptance

Proposing new features that impact all SDKs will occur by creating a Github Issue in this repo. The Issue should include motivation or rationale in addition to any relevant reading material. New features will be discussed and decided upon during weekly syncs

> [!IMPORTANT]
> Language agnostic test vectors **must** be produced _prior_ to commencing implementation beyond the first SDK

Test vectors are in the test-vector directory. More info on test vectors can be found there.

## Implementation Criteria

An individual SDK will consider a feature implemented once the following requirements have been met:

- Produce/update API reference documentation for any/all methods added to the _public_ API surface. documenting private API methods is optional
- Produce/update relevant example usage
- Test coverage must be provided for _all_ public API methods.
- Tests that consume shared test vectors must all be passing

## Review Criteria

- Propose API surface design _prior_ to implementation. This can be done by opening a feature request Issue and fostering a discussion around the proposed design. This can be followed by a corresponding draft PR and providing prototypal code.

- Approval Required from a minimum of 2 people.

> [!NOTE]
> Requiring two reviewers will likely slow things down at the benefit of ensuring there's a sufficient amount of visibility on the changes being made. It also prevents bottlenecks from forming. Many people on our team should be able to speak to the features landing in our SDKs.

## Release Criteria

- Each Release will have an accompanying Github Release
- Each Github Release will have an accompanying git tag that matches the version being published
- For the forseeable future, each SDK is free to publish releases at a frequency that the SDK's respective DRI sees fit

## CI / CD

Each SDK will use GitHub Actions for CI/CD and other automations.

Find the latest [Projects Health Dashboard here](https://developer.tbd.website/open-source/projects-dashboard/). 

## Publishing Artifacts

Each SDK will be published to the most widely adopted registry/repository for the respective language

| SDK        | Repository          |
| ---------- | ------------------- |
| Typescript | npm                 |
| Kotlin     | maven central       |
| Rust       | crates.io           |
| Swift      | swift package index |
| Go         | tbd                 |

## Publishing API Reference Documentation

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

## Example Feature Usage

Each SDK will **publish** example usage for _each_ implemented feature. This can either be included as a part of API reference documentation _or_ published separately.

## Test Vectors

Test vectors ensure interoporability of features across SDKs and language implementations by providing common test cases with an input and expected output pair. They include both success and failure cases that can be vectorized.

This repo serves as the home for all Web5 feature related vectors. They are available in the [test-vectors](./test-vectors/) directory.

The `sdk-report-runner` repo consumes the output tests for these test vectors in each repo and generates a report - [report-runner](https://github.com/TBD54566975/sdk-report-runner).

### Adding & Updating Test Vectors

New test vectors should follow the standard [vector structure](./test-vectors/). Vectors are automatically validated against their [JSON Schema](https://json-schema.org/) via CI.

Create a PR in this repo for adding / updating Web5 test vectors.