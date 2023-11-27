package tests

import (
	"context"
	"fmt"
	"strings"

	"github.com/TBD54566975/sdk-development/openapi"
)

func init() {
	tests["did:ion"] = map[string]testfn{
		"CreateRequest": didIonCreateRequest,
		// "UpdateRequest":     didIonUpdateRequest,
		// "RecoverRequest":    didIonRecoverRequest,
		// "DeactivateRequest": didIonDeactivateRequest,
		// "Resolution":        didIonResolution,
		// "Anchoring":         didIonAnchoring,
	}
}

func didIonCreateRequest(ctx context.Context, serverURL string) []error {
	client, err := openapi.NewClientWithResponses(serverURL)
	if err != nil {
		return []error{err}
	}

	response, err := client.DidIonCreateWithResponse(ctx)
	if err != nil {
		return []error{err}
	}

	if response.JSON200 == nil {
		return unexpectedResponseCode(response.HTTPResponse, response.Body)
	}

	didParts := strings.Split(response.JSON200.Did, ":")
	if len(didParts) != 4 {
		return []error{fmt.Errorf("invalid did:ion returned: 4 parts expected, %d found: %s", len(didParts), didParts)}
	}

	errs := []error{}
	if err := compareStrings(didParts[0], "did", "did 1st part"); err != nil {
		errs = append(errs, err)
	}

	if err := compareStrings(didParts[1], "ion", "did 2nd part"); err != nil {
		errs = append(errs, err)
	}

	if len(didParts[2]) != 46 {
		errs = append(errs, fmt.Errorf("3rd part of returned did of unexpected length: expected 46 characters, got %d", len(didParts[2])))
	}

	// TODO: base64decode didPart[3] and validate it's JSON structure

	return errs
}
