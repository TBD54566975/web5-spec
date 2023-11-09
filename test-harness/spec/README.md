# Test Harness
<!-- markdownlint-disable -->

## Dependencies

* [Encoders SDK](#encoders-sdk)
  * CBOR
  * SHA256
  * Base64
* [Crypto SDK](#crypto-sdk)
  * Key Manager
  * JWS Signatures
  * JWS Verification
* [DIDs SDK](#dids-sdk)
  * DID Resolution
* [Credentials SDK](#credentials-sdk)
  * Concrete Presentation Definition Types
  * Concrete Presentation Submission Types
  * Presentation Exchange Eval - Evaluate Presentation Submission against Presentation Definition
  * VC Verification
  * VC Revocation Check
    * StatusList Lookup
  * VP Verification

## Features

| Feature                   | Typescript | Kotlin | Swift |
| ------------------------- | ---------- | ------ | ----- |
| Create Offering           | :x:        | :x:    | :x:   |
| Create Resource           | :x:        | :x:    | :x:   |
| Verify Resource Integrity | :x:        | :x:    | :x:   |
| Create RFQ                | :x:        | :x:    | :x:   |
| Create Quote              | :x:        | :x:    | :x:   |
| Create Order              | :x:        | :x:    | :x:   |
| Create OrderStatus        | :x:        | :x:    | :x:   |
| Create Close              | :x:        | :x:    | :x:   |
| Create Message            | :x:        | :x:    | :x:   |
| Sign Message              | :x:        | :x:    | :x:   |
| Verify Message Integrity  | :x:        | :x:    | :x:   |
| Validate Message          | :x:        | :x:    | :x:   |
| Hash Private fields       | :x:        | :x:    | :x:   |

References:

* [tbDEX Protocol Spec](https://github.com/TBD54566975/tbdex-protocol/blob/rest-api-spec/README.md)

## DIDs SDK

SDK used to create and resolve DIDs

Dependencies:

* [Encoders SDK](#encoders-sdk)
* [Crypto SDK](#crypto-sdk)

## DID ION

| Feature                                                                      | Typescript     | Kotlin | Swift |
| ---------------------------------------------------------------------------- | -------------- | ------ | ----- |
| [`CreateRequest`](https://identity.foundation/sidetree/spec/#create)         | ✅              | :x:    | :x:   |
| [`UpdateRequest`](https://identity.foundation/sidetree/spec/#update)         | :construction: | :x:    | :x:   |
| [`RecoverRequest`](https://identity.foundation/sidetree/spec/#recover)       | :construction: | :x:    | :x:   |
| [`DeactivateRequest`](https://identity.foundation/sidetree/spec/#deactivate) | :construction: | :x:    | :x:   |
| Resolution                                                                   | ✅              | :x:    | :x:   |
| Anchoring                                                                    | ✅              | :x:    | :x:   |

* `CreateRequest` - used to create an ION DID
* `UpdateRequest` - used to update an ION DID
* `RecoverRequest` - used to recover an ION DID
* `DeactivateRequest` - used to decomission an ION DID
* Resolution - used to resolve an ION DID.
  * Needed for decentralized discovery
  * Needed for signature verification

## DID Key

| Feature    | Typescript | Kotlin | Swift |
| ---------- | ---------- | ------ | ----- |
| Creation   | ✅          | :x:    | :x:   |
| Resolution | ✅          | :x:    | :x:   |

References:

* [DID Spec](https://www.w3.org/TR/did-core/)
* [DID ION aka Sidetree Spec](https://identity.foundation/sidetree/spec/)
* [Typescript DIDs Package](https://github.com/TBD54566975/web5-js/tree/main/packages/dids)

# Credentials SDK

SDK used to create, verify and construct data structures needed to exchange Verifiable Credentials

## Dependencies

* [Encoders SDK](#encoders-sdk)
* [Crypto SDK](#crypto-sdk)
* [DIDs SDK](#dids-sdk)

## Features

| Feature               | Typescript | Kotlin         | Swift |
| --------------------- | ---------- | -------------- | ----- |
| Presentation Exchange | ✅          | :construction: | :x:   |
| VC Creation           | ✅          | :construction: | :x:   |
| VC Verification       | ✅          | :construction: | :x:   |
| VP Creation           | ✅          | :construction: | :x:   |
| VP Verification       | ✅          | :construction: | :x:   |
| StatusList lookup     | :x:        | :x:            | :x:   |

* [Presentation Exchange](https://identity.foundation/presentation-exchange/)
  * Presentation Definition - Data Structure used by PFIs to express their VC requirements
  * Presentation Submission - Data Structure used by Alice to fulfill the VC requirements expressed by a PFI
    * MVP must include filtering by issuer ID and type.

* VC Creation - Used by Issuer to create Verifiable Credentials
* VC Verification - Used by PFI to perform integrity checks on VCs provided by Alice
* VP Creation - Container used by Alice to pack the Presentation Submission created to fulfill the VC requirements expressed by a PFI
* VP Verification - Used by PFI to perform integrity check on Presentation Submission provided by Alice
* [StatusList](https://www.w3.org/TR/vc-status-list/) Lookup - Used by PFI as part of privacy preserving integrity check

References:

* [VC Spec](https://www.w3.org/TR/vc-data-model/)
* [Presentation Exchange Spec](https://identity.foundation/presentation-exchange/)
* [StatusList 2021 Spec](https://www.w3.org/TR/vc-status-list/)
* [Typescript Credentials Package](https://github.com/TBD54566975/web5-js/tree/main/packages/credentials)
  * Open PRs:
    * [PEX](https://github.com/TBD54566975/web5-js/pull/164)
    * [VC Creation](https://github.com/TBD54566975/web5-js/pull/148)

# Crypto SDK

SDK that contains the cryptographic foundation needed for DIDs, VCs, and tbDEX messages

## Key Generation

| Feature     | Typescript | Kotlin | Swift |
| ----------- | ---------- | ------ | ----- |
| `secp256k1` | ✅          | :x:    | :x:   |
| `Ed25519`   | ✅          | :x:    | :x:   |
| `secp256r1` | :x:        | :x:    | :x:   |

> :information_source: `secp256r1` is a FIPS compliant algorithm needed to support government use-cases. **Stretch Goal**

## Signing

| Feature     | Typescript | Kotlin | Swift |
| ----------- | ---------- | ------ | ----- |
| `secp256k1` | ✅          | :x:    | :x:   |
| `Ed25519`   | ✅          | :x:    | :x:   |
| `secp256r1` | :x:        | :x:    | :x:   |

## Verification

| Feature     | Typescript | Kotlin | Swift |
| ----------- | ---------- | ------ | ----- |
| `secp256k1` | ✅          | :x:    | :x:   |
| `Ed25519`   | ✅          | :x:    | :x:   |
| `secp256r1` | :x:        | :x:    | :x:   |

## JOSE

| Feature      | Typescript     | Kotlin | Swift |
| ------------ | -------------- | ------ | ----- |
| `JWS create` | :construction: | :x:    | :x:   |
| `JWS verify` | :construction: | :x:    | :x:   |
| `JWK encode` | ✅              | :x:    | :x:   |
| `JWK decode` | ✅              | :x:    | :x:   |
| `JWT create` | :construction: | :x:    | :x:   |
| `JWT verify` | :construction: | :x:    | :x:   |

> :information_source: prefer Compact JWS. Can work with General JWS

## Key Manager

| Feature        | Typescript | Kotlin | Swift |
| -------------- | ---------- | ------ | ----- |
| `Generate Key` | ✅          | :x:    | :x:   |
| `Import Key`   | ✅          | :x:    | :x:   |
| `Sign`         | ✅          | :x:    | :x:   |
| `Verify`       | ✅          | :x:    | :x:   |

References:

* [JWS Spec](https://datatracker.ietf.org/doc/html/rfc7515)
* [JWK Spec](https://datatracker.ietf.org/doc/html/rfc7517)
* [JWT Spec](https://datatracker.ietf.org/doc/html/rfc7519)
* [Typescript Crypto Package](https://github.com/TBD54566975/web5-js/tree/main/packages/crypto)

# Encoders SDK

SDK that contains encode/decode utility functions needed for DIDs, VCs, and tbDEX

| Feature         | Typescript | Kotlin | Swift |
| --------------- | ---------- | ------ | ----- |
| `base64 encode` | ✅          | :x:    | :x:   |
| `base64 decode` | ✅          | :x:    | :x:   |
| `base58 encode` | ✅          | :x:    | :x:   |
| `base58 decode` | ✅          | :x:    | :x:   |
| `sha256 encode` | ✅          | :x:    | :x:   |
| `cbor encode`   | :x:        | :x:    | :x:   |
| `cbor decode`   | :x:        | :x:    | :x:   |

References:

* [Typescript Encoders Package](https://github.com/TBD54566975/web5-js/tree/main/packages/common)
