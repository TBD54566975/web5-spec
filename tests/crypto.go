package tests

import (
	"context"

	"github.com/TBD54566975/web5-spec/openapi"
)

func init() {
	tests["Crypto SDK"] = map[string]testfn{
		"Generate secp256k1 key": generateSecp256k1Key,
		"Generate secp256r1 key": generateSecp256r1Key,
		"Generate ed25519 key":   generateEd25519Key,
	}
}

func generateSecp256k1Key(ctx context.Context, serverURL string) []error {
	client, err := openapi.NewClientWithResponses(serverURL)
	if err != nil {
		return []error{err}
	}

	response, err := client.CryptoGenerateKeySecp256k1WithResponse(ctx)
	if err != nil {
		return []error{err}
	}

	if response.JSON200 == nil {
		return unexpectedResponseCode(response.HTTPResponse, response.Body)
	}

	return nil
}

func generateEd25519Key(ctx context.Context, serverURL string) []error {
	client, err := openapi.NewClientWithResponses(serverURL)
	if err != nil {
		return []error{err}
	}

	response, err := client.CryptoGenerateKeyEd25519WithResponse(ctx)
	if err != nil {
		return []error{err}
	}

	if response.JSON200 == nil {
		return unexpectedResponseCode(response.HTTPResponse, response.Body)
	}

	return nil
}

func generateSecp256r1Key(ctx context.Context, serverURL string) []error {
	client, err := openapi.NewClientWithResponses(serverURL)
	if err != nil {
		return []error{err}
	}

	response, err := client.CryptoGenerateKeySecp256r1WithResponse(ctx)
	if err != nil {
		return []error{err}
	}

	if response.JSON200 == nil {
		return unexpectedResponseCode(response.HTTPResponse, response.Body)
	}

	return nil
}
