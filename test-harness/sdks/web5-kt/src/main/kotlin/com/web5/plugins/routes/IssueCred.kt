package com.web5.plugins.routes

import com.danubetech.verifiablecredentials.CredentialSubject
import com.web5.models.CredentialIssuanceRequest
import com.web5.models.CredentialIssuanceResponse
import com.web5.models.StringEncodedData
import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import web5.sdk.credentials.VcDataModel
import web5.sdk.credentials.VerifiableCredential
import web5.sdk.crypto.InMemoryKeyManager
import web5.sdk.dids.methods.key.DidKey
import java.net.URI
import java.util.*

suspend fun ApplicationCall.credentialIssue() {
    val payload = receive<CredentialIssuanceRequest>()
    val reqVc = payload.credential

    val keyManager = InMemoryKeyManager()
    val issuerDid = DidKey.create(keyManager)

    val vc = VerifiableCredential.create(reqVc.type[reqVc.type.size - 1], reqVc.issuer, reqVc.credentialSubject.get("id") as String, reqVc.credentialSubject)

    val vcJwt = vc.sign(issuerDid)
    val res = CredentialIssuanceResponse(verifiableCredential = StringEncodedData(data = vcJwt))

    respond(HttpStatusCode.OK, res)
}