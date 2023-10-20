import { Request, Response } from "express";
import { LocalKms, KeyManager, Web5ManagedAgent } from "@web5/agent";
import { paths } from "./openapi.js";
import { Convert } from "@web5/common";

// @ts-ignore
const mockAgent: Web5ManagedAgent = {};
const localKms = new LocalKms({ kmsName: "memory", agent: mockAgent });
const keyManager = new KeyManager({
  kms: { memory: localKms },
  agent: mockAgent,
});

export async function cryptoSecp256k1Generate(req: Request, res: Response) {
  const key = await keyManager.generateKey({
    kms: "memory",
    algorithm: { name: "ECDSA", namedCurve: "secp256k1" },
    keyUsages: ["sign", "verify"],
  });

  const resp: paths["/crypto/generate-key/secp256k1"]["post"]["responses"]["200"]["content"]["application/json"] =
    {
      public: Convert.uint8Array(key.publicKey.material!).toBase64Url(),
      private: null, //Convert.uint8Array(key.privateKey.material!).toBase64Url(), // private key material here is null?
    };

  res.json(resp);
}

export async function cryptoEd25519Generate(req: Request, res: Response) {
  const keyPair = await keyManager.generateKey({
    kms: "memory",
    algorithm: { name: "EdDSA", namedCurve: "Ed25519" },
    keyUsages: ["sign", "verify"],
  });

  const resp: paths["/crypto/generate-key/ed25519"]["post"]["responses"]["200"]["content"]["application/json"] =
    {
      public: Convert.uint8Array(keyPair.publicKey.material!).toBase64Url(),
      private: null, //Convert.uint8Array(keyPair.privateKey.material!).toBase64Url(),
    };

  res.json(resp);
}
