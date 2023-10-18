import { Request, Response } from "express";
import { VerifiableCredentialTypeV1 } from "@web5/credentials";
import { paths } from "./openapi.js";

export async function credentialIssue(req: Request, res: Response) {
  const body: paths["/credentials/issue"]["post"]["requestBody"]["content"]["application/json"] =
  req.body;
  
  const vc:VerifiableCredentialTypeV1 = body.credential;
  res.json({verifiableCredential:vc});
}