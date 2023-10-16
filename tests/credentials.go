package tests

import (
	"context"

	"github.com/TBD54566975/web5-spec/openapi"
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

func vcCreate(ctx context.Context, serverURL string) []error {
	expectedContext := []string{"https://www.w3.org/2018/credentials/v1"}
	expectedType := []string{"VerifiableCredential"}
	expectedID := "id-123"
	expectedIssuer := "did:example:123"
	expectedCredentialSubject := map[string]interface{}{
		"id":        "did:example:123",
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
			Issuer: openapi.CredentialIssuer{Id: expectedIssuer},
			Type:   expectedType,
		},
	})
	if err != nil {
		return []error{err}
	}

	if response.JSON200 == nil {
		return unexpectedResponseCode(response.HTTPResponse, response.Body)
	}

	vc := response.JSON200.VerifiableCredential

	// Check @context
	errs := []error{}
	if err := compareStringSlices(vc.Context, expectedContext, "@context"); err != nil {
		errs = append(errs, err)
	}

	// Check credentialSubject
	if err := compareMaps(vc.CredentialSubject, expectedCredentialSubject, "credentialSubject"); err != nil {
		errs = append(errs, err)
	}

	// Check id
	if err := compareStrings(vc.Id, expectedID, "id"); err != nil {
		errs = append(errs, err)
	}

	// Check type
	if err := compareStringSlices(vc.Type, expectedType, "type"); err != nil {
		errs = append(errs, err)
	}

	// Check issuer
	if err := compareStrings(vc.Issuer.Id, expectedIssuer, "issuer.id"); err != nil {
		errs = append(errs, err)
	}

	return errs
}
