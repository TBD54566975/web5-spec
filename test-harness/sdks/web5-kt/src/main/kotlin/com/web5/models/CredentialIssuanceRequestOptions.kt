/**
* web5 SDK test server
* No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
*
* The version of the OpenAPI document: 0.2.0
* 
*
* NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
* https://openapi-generator.tech
* Do not edit the class manually.
*/
package com.web5.models

/**
 * 
 * @param created 
 * @param challenge 
 * @param domain 
 * @param credentialStatus 
 */
data class CredentialIssuanceRequestOptions(
    val created: kotlin.String,
    val challenge: kotlin.String,
    val domain: kotlin.String,
    val credentialStatus: CredentialStatus
) 
