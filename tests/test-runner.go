package tests

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/exp/slog"
)

type testfn func(ctx context.Context, serverURL string) []error

var tests = map[string]map[string]testfn{}

func RunTests(serverURL string) map[string]map[string][]error {
	results := map[string]map[string][]error{}
	for group, ts := range tests {
		slog.Info("starting test group", "group", group)
		results[group] = map[string][]error{}
		for t, fn := range ts {
			slog.Info("running", "test", t)
			ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
			defer cancel()
			results[group][t] = fn(ctx, serverURL)
			if results[t] != nil {
				slog.Error("error", "test", t, "error", results[t])
			}
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

func unexpectedResponseCode(r *http.Response, body []byte) []error {
	sanatizedBody := strings.ReplaceAll(string(body), "\n", "\\n")
	return []error{fmt.Errorf("%s: %s", r.Status, sanatizedBody)}
}

func generateRandomString(len int) string {
	return "TODO: MAKE THIS RANDOM"
}
