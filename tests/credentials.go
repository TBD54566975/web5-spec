package tests

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/TBD54566975/web5-spec/openapi"
	"gopkg.in/square/go-jose.v2"
)

func init() {
	tests["Credentials SDK"] = map[string]testfn{
		"VC Create": vcCreate,
		// "VC Verify": vcVerify,
		// "VP Create": vpCreate,
		// "VP Verify": vpVerify,
		// "StatusList Lookup": statusListLookup,
	}
}

type Payload struct {
	Iss string                                      `json:"iss"`
	Sub string                                      `json:"sub"`
	Vc  openapi.CredentialIssuanceRequestCredential `json:"vc"`
}

func vcCreate(ctx context.Context, serverURL string) []error {
	expectedContext := []string{"https://www.w3.org/2018/credentials/v1"}
	expectedType := []string{"VerifiableCredential"}
	expectedID := "did:example:321"
	expectedIssuer := "did:example:123"

	expectedCredentialSubjectId := "did:example:123"
	expectedCredentialSubject := map[string]interface{}{
		"id":        expectedCredentialSubjectId,
		"firstName": "bob",
	}

	client, err := openapi.NewClientWithResponses(serverURL)
	if err != nil {
		return []error{err}
	}

	response, err := client.CredentialIssueWithResponse(ctx, openapi.CredentialIssuanceRequest{
		Credential: openapi.CredentialIssuanceRequestCredential{
			Context:           expectedContext,
			CredentialSubject: expectedCredentialSubject,
			// ExpirationDate: ,
			Id: expectedID,
			// IssuanceDate: ,
			Issuer: expectedIssuer,
			Type:   expectedType,
		},
	})

	if err != nil {
		return []error{err}
	}

	if response.JSON200 == nil {
		return unexpectedResponseCode(response.HTTPResponse, response.Body)
	}

	vcJwt := response.JSON200.VerifiableCredential.Data

	token, err := jose.ParseSigned(vcJwt)
	if err != nil {
		return []error{err}
	}
	payloadBytes := token.UnsafePayloadWithoutVerification()

	var payload Payload
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return []error{fmt.Errorf("failed to unmarshal payload into struct: %v", err)}
	}

	errs := []error{}

	// Check iss
	if err := compareStringsContains(payload.Iss, "did:", "iss"); err != nil {
		errs = append(errs, err)
	}

	// Check sub
	if err := compareStrings(payload.Sub, expectedCredentialSubjectId, "sub"); err != nil {
		errs = append(errs, err)
	}

	// Check @context
	if err := compareStringSlices(payload.Vc.Context, expectedContext, "@context"); err != nil {
		errs = append(errs, err)
	}

	// Check credentialSubject
	if err := compareMaps(payload.Vc.CredentialSubject, expectedCredentialSubject, "credentialSubject"); err != nil {
		errs = append(errs, err)
	}

	// Check id
	if err := compareStringsContains(payload.Vc.Id, "did:", "id"); err != nil {
		errs = append(errs, err)
	}

	// Check type
	if err := compareStringSlices(payload.Vc.Type, expectedType, "type"); err != nil {
		errs = append(errs, err)
	}

	if err := compareStrings(payload.Vc.Issuer, expectedIssuer, "issuer"); err != nil {
		errs = append(errs, err)
	}

	//switch issuer := payload.Vc.Issuer.(type) {
	//case string:
	//	if err := compareStrings(issuer, expectedIssuer, "issuer"); err != nil {
	//		errs = append(errs, err)
	//	}
	//case openapi.CredentialIssuer:
	//	if err := compareStrings(issuer.Id, expectedIssuer, "issuer.id"); err != nil {
	//		errs = append(errs, err)
	//	}
	//default:
	//	errs = append(errs, fmt.Errorf("Unexpected type for Issuer"))
	//}

	return errs
}
