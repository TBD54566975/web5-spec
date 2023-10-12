package tests

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/exp/slog"
)

type test struct {
	Name string
	Fn   func(ctx context.Context, serverURL string) error
}

var tests = []test{
	{Name: "CredentialIssuance", Fn: CredentialIssuanceTest},
}

func RunTests(serverURL string) map[string]error {
	results := map[string]error{}
	for _, t := range tests {
		slog.Info("running", "test", t.Name)
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		results[t.Name] = t.Fn(ctx, serverURL)
		if results[t.Name] != nil {
			slog.Error("error", "test", t.Name, "error", results[t.Name])
		}
	}

	return results
}

func compareStringSlices(actual []string, expected []string, field string) error {
	if len(actual) != len(expected) {
		return fmt.Errorf("incorrect value for %s: expected %v, got %v", field, expected, actual)
	}

	for i, v := range actual {
		if v != expected[i] {
			return fmt.Errorf("incorrect value for %s: expected %v, got %v", field, expected, actual)
		}
	}
	return nil
}

// Utility function to compare maps
func compareMaps(actual, expected map[string]interface{}, field string) error {
	for k, v := range expected {
		if actual[k] != v {
			return fmt.Errorf("incorrect value for %s: expected %v, got %v", field, expected, actual)
		}
	}
	return nil
}

func compareStrings(actual string, expected string, field string) error {
	if actual != expected {
		return fmt.Errorf("incorrect value for %s: expected %v, got %v", field, expected, actual)
	}
	return nil
}
