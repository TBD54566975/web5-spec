package tests

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/TBD54566975/sdk-development/openapi"
	"github.com/mr-tron/base58"
)

func init() {
	tests["encoders"] = map[string]testfn{
		"base64 encode": encoderBase64Encode,
		"base64 decode": encoderBase64Decode,
		"base58 encode": encoderBase58Encode,
		"base58 decode": encoderBase58Decode,
		// "CBOR encode":   encoderCBOREncode,
		// "CBOR decode":   encoderCBORDecode,
		"sha256 encode": encoderSha256Encode,
	}
}

func encoderBase64Decode(ctx context.Context, serverURL string) []error {
	client, err := openapi.NewClientWithResponses(serverURL)
	if err != nil {
		return []error{err}
	}

	expected := generateRandomString(128)
	b64 := base64.RawURLEncoding.EncodeToString([]byte(expected))

	resp, err := client.EncodersBase64DecodeWithResponse(ctx, openapi.StringEncodedData{Data: b64})
	if err != nil {
		return []error{err}
	}

	if resp.JSON200 == nil {
		return unexpectedResponseCode(resp.HTTPResponse, resp.Body)
	}

	errs := []error{}
	if err := compareStrings(resp.JSON200.Data, expected, "base64 decode"); err != nil {
		errs = append(errs, err)
	}

	return errs
}

func encoderBase58Encode(ctx context.Context, serverURL string) []error {
	client, err := openapi.NewClientWithResponses(serverURL)
	if err != nil {
		return []error{err}
	}

	input := generateRandomString(128)
	expected := base58.Encode([]byte(input))

	resp, err := client.EncodersBase58EncodeWithResponse(ctx, openapi.StringEncodedData{Data: input})
	if err != nil {
		return []error{err}
	}

	if resp.JSON200 == nil {
		return unexpectedResponseCode(resp.HTTPResponse, resp.Body)
	}

	errs := []error{}
	if err := compareStrings(resp.JSON200.Data, expected, "base58 encode"); err != nil {
		errs = append(errs, err)
	}

	return errs
}

func encoderBase58Decode(ctx context.Context, serverURL string) []error {
	client, err := openapi.NewClientWithResponses(serverURL)
	if err != nil {
		return []error{err}
	}

	expected := generateRandomString(128)

	b58 := base58.Encode([]byte(expected))

	resp, err := client.EncodersBase58DecodeWithResponse(ctx, openapi.StringEncodedData{Data: b58})
	if err != nil {
		return []error{err}
	}

	if resp.JSON200 == nil {
		return unexpectedResponseCode(resp.HTTPResponse, resp.Body)
	}

	errs := []error{}
	if err := compareStrings(resp.JSON200.Data, expected, "base58 decode"); err != nil {
		errs = append(errs, err)
	}

	return errs
}

func encoderBase64Encode(ctx context.Context, serverURL string) []error {
	client, err := openapi.NewClientWithResponses(serverURL)
	if err != nil {
		return []error{err}
	}

	input := generateRandomString(128)
	expected := base64.RawURLEncoding.EncodeToString([]byte(input))

	resp, err := client.EncodersBase64EncodeWithResponse(ctx, openapi.StringEncodedData{Data: input})
	if err != nil {
		return []error{err}
	}

	if resp.JSON200 == nil {
		return unexpectedResponseCode(resp.HTTPResponse, resp.Body)
	}

	errs := []error{}
	if err := compareStrings(resp.JSON200.Data, expected, "base64 encode"); err != nil {
		errs = append(errs, err)
	}

	return errs
}

func encoderSha256Encode(ctx context.Context, serverURL string) []error {
	client, err := openapi.NewClientWithResponses(serverURL)
	if err != nil {
		return []error{err}
	}

	input := generateRandomString(128)
	expected := fmt.Sprintf("%x", sha256.Sum256([]byte(input)))

	resp, err := client.EncodersSha256EncodeWithResponse(ctx, openapi.StringEncodedData{Data: input})
	if err != nil {
		return []error{err}
	}

	if resp.JSON200 == nil {
		return unexpectedResponseCode(resp.HTTPResponse, resp.Body)
	}

	errs := []error{}
	if err := compareStrings(resp.JSON200.Data, expected, "sha256 encode"); err != nil {
		errs = append(errs, err)
	}

	return errs
}
