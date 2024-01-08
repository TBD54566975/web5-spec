package reports

import (
	"archive/zip"
	"bytes"
	"fmt"
	"strings"

	junit "github.com/joshdk/go-junit"
	"golang.org/x/exp/slog"
)

func readArtifactZip(artifact []byte) ([]junit.Suite, error) {
	z, err := zip.NewReader(bytes.NewReader(artifact), int64(len(artifact)))
	if err != nil {
		return nil, err
	}

	suites := []junit.Suite{}
	for _, f := range z.File {
		if !strings.HasSuffix(f.Name, ".xml") {
			continue
		}

		s, err := readJunit(f)
		if err != nil {
			return nil, fmt.Errorf("error reading %s: %v", f.Name, err)
		}

		slog.Info("read", "suites", len(s), "file", f.Name)

		suites = append(suites, s...)
	}

	return suites, nil
}

func readJunit(f *zip.File) ([]junit.Suite, error) {
	r, err := f.Open()
	if err != nil {
		return []junit.Suite{}, err
	}
	defer r.Close()

	suites, err := junit.IngestReader(r)
	if err != nil {
		return []junit.Suite{}, err
	}

	return suites, nil
}
