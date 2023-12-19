package reports

import (
	"io/fs"
	"path/filepath"
	"strings"
)

var knownVectors = readKnownVectors()

func readKnownVectors() map[string]map[string]bool {
	knownVectors := make(map[string]map[string]bool)
	dir := "../web5-test-vectors"
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".json") || strings.HasSuffix(path, ".schema.json") {
			return nil
		}

		feature, vector := parseVectorPath(strings.TrimPrefix(path, dir))
		if knownVectors[feature] == nil {
			knownVectors[feature] = make(map[string]bool)
		}
		knownVectors[feature][vector] = true

		return nil
	})

	if err != nil {
		panic(err)
	}

	return knownVectors
}

func parseVectorPath(path string) (feature string, vector string) {
	feature, vector = filepath.Split(path)
	vector = strings.TrimSuffix(vector, ".json")
	feature = strings.Trim(feature, "/")

	featureWords := []string{}
	for _, word := range strings.Split(feature, "_") {
		featureWords = append(featureWords, strings.Title(word)) // TODO: strings.Title is deprecated
	}

	feature = strings.Join(featureWords, "")

	return feature, vector
}
