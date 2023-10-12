# web5-spec

A compliance suite for web5 SDKs

## Dependencies

- [Encoders SDK](#Encoders-SDK)
  - CBOR
  - SHA256
  - Base64
- [Crypto SDK](#Crypto-SDK)
  - Key Manager
  - JWS Signatures
  - JWS Verification
- [DIDs SDK](#DIDs-SDK)
  - DID Resolution
- [Credentials SDK](#Credentials-SDK)
  - Concrete Presentation Definition Types
  - Concrete Presentation Submission Types
  - Presentation Exchange Eval - Evaluate Presentation Submission against Presentation Definition
  - VC Verification
  - VC Revocation Check
    - StatusList Lookup
  - VP Verification

## Features

| Feature                   | Typescript | Kotlin | Swift |
| ------------------------- | ---------- | ------ | ----- |
| Create Offering           | ❌         | ❌     | ❌    |
| Create Resource           | ❌         | ❌     | ❌    |
| Verify Resource Integrity | ❌         | ❌     | ❌    |
| Create RFQ                | ❌         | ❌     | ❌    |
| Create Quote              | ❌         | ❌     | ❌    |
| Create Order              | ❌         | ❌     | ❌    |
| Create OrderStatus        | ❌         | ❌     | ❌    |
| Create Close              | ❌         | ❌     | ❌    |
| Create Message            | ❌         | ❌     | ❌    |
| Sign Message              | ❌         | ❌     | ❌    |
| Verify Message Integrity  | ❌         | ❌     | ❌    |
| Validate Message          | ❌         | ❌     | ❌    |
| Hash Private fields       | ❌         | ❌     | ❌    |

## References

- [tbDEX Protocol Spec](https://github.com/TBD54566975/tbdex-protocol/blob/rest-api-spec/README.md)

# DIDs SDK

SDK used to create and resolve DIDs

## Dependencies

- [Encoders SDK](#Encoders-SDK)
- [Crypto SDK](#Crypto-SDK)

## DID ION

### Features

| Feature                                                                      | Typescript | Kotlin | Swift |
| ---------------------------------------------------------------------------- | ---------- | ------ | ----- |
| [`CreateRequest`](https://identity.foundation/sidetree/spec/#create)         | ✅         | ❌     | ❌    |
| [`UpdateRequest`](https://identity.foundation/sidetree/spec/#update)         | 🚧         | ❌     | ❌    |
| [`RecoverRequest`](https://identity.foundation/sidetree/spec/#recover)       | 🚧         | ❌     | ❌    |
| [`DeactivateRequest`](https://identity.foundation/sidetree/spec/#deactivate) | 🚧         | ❌     | ❌    |
| Resolution                                                                   | ✅         | ❌     | ❌    |
| Anchoring                                                                    | ✅         | ❌     | ❌    |

- `CreateRequest` - used to create an ION DID
- `UpdateRequest` - used to update an ION DID
- `RecoverRequest` - used to recover an ION DID
- `DeactivateRequest` - used to decomission an ION DID
- Resolution - used to resolve an ION DID.
  - Needed for decentralized discovery
  - Needed for signature verification

## DID Key

### Features

| Feature    | Typescript | Kotlin | Swift |
| ---------- | ---------- | ------ | ----- |
| Creation   | ✅         | ❌     | ❌    |
| Resolution | ✅         | ❌     | ❌    |

## References

- [DID Spec](https://www.w3.org/TR/did-core/)
- [DID ION aka Sidetree Spec](https://identity.foundation/sidetree/spec/)
- [Typescript DIDs Package](https://github.com/TBD54566975/web5-js/tree/main/packages/dids)

# Credentials SDK

SDK used to create, verify and construct data structures needed to exchange Verifiable Credentials

## Dependencies

- [Encoders SDK](#Encoders-SDK)
- [Crypto SDK](#Crypto-SDK)
- [DIDs SDK](#DIDs-SDK)

## Features

| Feature               | Typescript | Kotlin | Swift |
| --------------------- | ---------- | ------ | ----- |
| Presentation Exchange | ✅         | 🚧     | ❌    |
| VC Creation           | ✅         | 🚧     | ❌    |
| VC Verification       | ✅         | 🚧     | ❌    |
| VP Creation           | ✅         | 🚧     | ❌    |
| VP Verification       | ✅         | 🚧     | ❌    |
| StatusList lookup     | ❌         | ❌     | ❌    |

- [Presentation Exchange](https://identity.foundation/presentation-exchange/)

  - Presentation Definition - Data Structure used by PFIs to express their VC requirements
  - Presentation Submission - Data Structure used by Alice to fulfill the VC requirements expressed by a PFI
    - MVP must include filtering by issuer ID and type.

- VC Creation - Used by Issuer to create Verifiable Credentials
- VC Verification - Used by PFI to perform integrity checks on VCs provided by Alice
- VP Creation - Container used by Alice to pack the Presentation Submission created to fulfill the VC requirements expressed by a PFI
- VP Verification - Used by PFI to perform integrity check on Presentation Submission provided by Alice
- [StatusList](https://www.w3.org/TR/vc-status-list/) Lookup - Used by PFI as part of privacy preserving integrity check

## References

- [VC Spec](https://www.w3.org/TR/vc-data-model/)
- [Presentation Exchange Spec](https://identity.foundation/presentation-exchange/)
- [StatusList 2021 Spec](https://www.w3.org/TR/vc-status-list/)
- [Typescript Credentials Package](https://github.com/TBD54566975/web5-js/tree/main/packages/credentials)
  - Open PRs:
    - [PEX](https://github.com/TBD54566975/web5-js/pull/164)
    - [VC Creation](https://github.com/TBD54566975/web5-js/pull/148)

# Crypto SDK

SDK that contains the cryptographic foundation needed for DIDs, VCs, and tbDEX messages

## Key Generation

### Features

| Feature     | Typescript | Kotlin | Swift |
| ----------- | ---------- | ------ | ----- |
| `secp256k1` | ✅         | ❌     | ❌    |
| `Ed25519`   | ✅         | ❌     | ❌    |
| `secp256r1` | ❌         | ❌     | ❌    |

> ℹ️ `secp256r1` is a FIPS compliant algorithm needed to support government use-cases. **Stretch Goal**

## Signing

### Features

| Feature     | Typescript | Kotlin | Swift |
| ----------- | ---------- | ------ | ----- |
| `secp256k1` | ✅         | ❌     | ❌    |
| `Ed25519`   | ✅         | ❌     | ❌    |
| `secp256r1` | ❌         | ❌     | ❌    |

## Verification

### Features

| Feature     | Typescript | Kotlin | Swift |
| ----------- | ---------- | ------ | ----- |
| `secp256k1` | ✅         | ❌     | ❌    |
| `Ed25519`   | ✅         | ❌     | ❌    |
| `secp256r1` | ❌         | ❌     | ❌    |

## JOSE

### Features

| Feature      | Typescript | Kotlin | Swift |
| ------------ | ---------- | ------ | ----- |
| `JWS create` | 🚧         | ❌     | ❌    |
| `JWS verify` | 🚧         | ❌     | ❌    |
| `JWK encode` | ✅         | ❌     | ❌    |
| `JWK decode` | ✅         | ❌     | ❌    |
| `JWT create` | 🚧         | ❌     | ❌    |
| `JWT verify` | 🚧         | ❌     | ❌    |

> ℹ️ prefer Compact JWS. Can work with General JWS

## Key Manager

### Features

| Feature        | Typescript | Kotlin | Swift |
| -------------- | ---------- | ------ | ----- |
| `Generate Key` | ✅         | ❌     | ❌    |
| `Import Key`   | ✅         | ❌     | ❌    |
| `Sign`         | ✅         | ❌     | ❌    |
| `Verify`       | ✅         | ❌     | ❌    |

## References

- [JWS Spec](https://datatracker.ietf.org/doc/html/rfc7515)
- [JWK Spec](https://datatracker.ietf.org/doc/html/rfc7517)
- [JWT Spec](https://datatracker.ietf.org/doc/html/rfc7519)
- [Typescript Crypto Package](https://github.com/TBD54566975/web5-js/tree/main/packages/crypto)

# Encoders SDK

SDK that contains encode/decode utility functions needed for DIDs, VCs, and tbDEX

## Features

| Feature         | Typescript | Kotlin | Swift |
| --------------- | ---------- | ------ | ----- |
| `base64 encode` | ✅         | ❌     | ❌    |
| `base64 decode` | ✅         | ❌     | ❌    |
| `base58 encode` | ✅         | ❌     | ❌    |
| `base58 decode` | ✅         | ❌     | ❌    |
| `sha256 encode` | ✅         | ❌     | ❌    |
| `cbor encode`   | ❌         | ❌     | ❌    |
| `cbor decode`   | ❌         | ❌     | ❌    |

## References

- [Typescript Encoders Package](https://github.com/TBD54566975/web5-js/tree/main/packages/common)
