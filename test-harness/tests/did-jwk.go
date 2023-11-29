package tests

import (
	"context"

	"github.com/TBD54566975/sdk-development/openapi"
)

func init() {
	tests["did:jwk"] = map[string]testfn{
		"Resolve": didJwkResolve,
	}
}

const (
	sampleDidJwk = "did:jwk:eyJjcnYiOiJQLTI1NiIsImt0eSI6IkVDIiwieCI6ImFjYklRaXVNczNpOF91c3pFakoydHBUdFJNNEVVM3l6OTFQSDZDZEgyVjAiLCJ5IjoiX0tjeUxqOXZXTXB0bm1LdG00NkdxRHo4d2Y3NEk1TEtncmwyR3pIM25TRSJ9"
)

func didJwkResolve(ctx context.Context, serverURL string) []error {
	client, err := openapi.NewClientWithResponses(serverURL)
	if err != nil {
		return []error{err}
	}

	response, err := client.DidJwkResolveWithResponse(ctx, openapi.DidResolutionRequest{Did: sampleDidJwk})
	if err != nil {
		return []error{err}
	}

	if response.JSON200 == nil {
		return unexpectedResponseCode(response.HTTPResponse, response.Body)
	}

	result := response.JSON200

	errs := []error{}
	if err := compareStrings(result.DidDocument.Id, sampleDidJwk, "the ID of the resolved did (should match the input)"); err != nil {
		errs = append(errs, err)
	}

	return errs
}
