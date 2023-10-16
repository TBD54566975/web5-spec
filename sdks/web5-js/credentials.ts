import { Request, Response } from 'express';
import { VcJwt, VerifiableCredential, SignOptions } from '@web5/credentials';
import { DidKeyMethod, PortableDid } from '@web5/dids';
import { Ed25519, Jose } from '@web5/crypto';
import { paths } from './openapi.js';

type Signer = (data: Uint8Array) => Promise<Uint8Array>;

let _ownDid: PortableDid;

async function getOwnDid(): Promise<PortableDid> {
    if(_ownDid) {
        return _ownDid;
    }
    _ownDid = await DidKeyMethod.create();
    return _ownDid;
}

export async function issueCredential(req: Request, res: Response) {
    const body: paths["/credentials/issue"]["post"]["requestBody"]["content"]["application/json"] = req.body;

    const ownDid = await getOwnDid()

    // build signing options
    const [signingKeyPair] = ownDid.keySet.verificationMethodKeys!;
    const privateKey = (await Jose.jwkToKey({ key: signingKeyPair.privateKeyJwk!})).keyMaterial;
    const subjectIssuerDid = ownDid.did;
    const signer = EdDsaSigner(privateKey);
    const signOptions: SignOptions = {
        issuerDid  : ownDid.did,
        subjectDid : ownDid.did,
        kid        : '#' + ownDid.did.split(':')[2],
        signer     : signer
    };

    const vcJwt: VcJwt = await VerifiableCredential.create(signOptions);
    // const resp: paths["/credentials/issue"]["post"]["responses"]["200"]["content"]["application/json"] = {
    //     verifiableCredential: {
    //     },
    // }
    res.json(vcJwt);
}

function EdDsaSigner(privateKey: Uint8Array): Signer {
    return async (data: Uint8Array): Promise<Uint8Array> => {
      const signature = await Ed25519.sign({ data, key: privateKey});
      return signature;
    };
  }