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
import web5.sdk.dids.DidKey
import java.net.URI
import java.util.*

suspend fun ApplicationCall.credentialIssue() {
    val payload = receive<CredentialIssuanceRequest>()
    val reqVc = payload.credential

    val keyManager = InMemoryKeyManager()
    val issuerDid = DidKey.create(keyManager)

    val credentialSubject = CredentialSubject.builder()
        .id(URI.create(reqVc.credentialSubject.get("id").toString()))
        .claims(reqVc.credentialSubject)
        .build()

    val vcDataModel = VcDataModel.builder()
        .id(URI.create(reqVc.id))
        .issuer(URI.create(reqVc.issuer))
        .issuanceDate(Date())
        .credentialSubject(credentialSubject)
        .build()

    val vc = VerifiableCredential(vcDataModel)

    val vcJwt = vc.sign(issuerDid)
    val res = CredentialIssuanceResponse(verifiableCredential = StringEncodedData(data = vcJwt))

    respond(HttpStatusCode.OK, res)
}