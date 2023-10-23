package com.web5.plugins.routes

import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import models.CredentialIssuanceRequestCredential
import models.CredentialIssuanceResponse
import models.StringEncodedData
import web5.sdk.credentials.VerifiableCredential
import web5.sdk.crypto.InMemoryKeyManager
import web5.sdk.dids.DidKey

suspend fun ApplicationCall.credentialIssue() {
    val payload = receive<CredentialIssuanceRequestCredential>()

    val keyManager = InMemoryKeyManager()
    val issuerDid = DidKey.create(keyManager)

    val vc = VerifiableCredential.create(
        type = payload.type[1],
        issuer = payload.issuer.id,
        subject = payload.credentialSubject.get("id").toString(),
        data = payload.credentialSubject
    )

    println(vc)

    val vcJwt = vc.sign(issuerDid)
    val res = CredentialIssuanceResponse(verifiableCredential = StringEncodedData(vcJwt))

    respond(HttpStatusCode.OK, res)
}