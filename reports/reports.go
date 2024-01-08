package reports

import (
	"embed"
	"errors"
	htmltemplate "html/template"
	"regexp"
	"time"

	junit "github.com/joshdk/go-junit"
	"golang.org/x/exp/slog"
)

var (
	ErrNotSupported = errors.New("test not supported by this SDK")

	//go:embed report-template.html
	templatesFS embed.FS

	htmlTemplates = htmltemplate.New("")
	funcmap       = map[string]any{
		"sanatizeHTML": sanatizeHTML,
	}
)

func init() {
	htmlTemplates.Funcs(funcmap)
	if _, err := htmlTemplates.ParseFS(templatesFS, "report-template.html"); err != nil {
		panic(err)
	}
}

type SDKMeta struct {
	Name         string
	Repo         string
	ArtifactName string
	FeatureRegex *regexp.Regexp
	VectorRegex  *regexp.Regexp
	VectorPath   string
}

type Report struct {
	SDK     SDKMeta
	Results map[string]map[string]Result
}

type Result struct {
	Exists bool
	Errors []error
	Time   time.Duration
}

func (r Report) IsPassing() bool {
	for _, results := range r.Results {
		for _, result := range results {
			if result.IsSkipped() {
				continue
			}

			if len(result.Errors) > 0 {
				return false
			}
		}

	}

	return true
}

func (r Result) IsSkipped() bool {
	return len(r.Errors) == 1 && r.Errors[0] == ErrNotSupported
}

func (r Result) GetEmoji() string {
	if !r.Exists {
		return "🚧"
	}

	if len(r.Errors) == 0 {
		return "✅"
	}

	return "❌"
}

func (r Result) GetEmojiAriaLabel() string {
	if !r.Exists {
		return "In progress"
	}

	if len(r.Errors) == 0 {
		return "Success"
	}

	return "Failed"
}

func (s SDKMeta) buildReport(suites []junit.Suite) (Report, error) {
	results := make(map[string]map[string]Result)

	for feature, vectors := range knownVectors {
		results[feature] = make(map[string]Result)
		for vector := range vectors {
			results[feature][vector] = Result{}
		}
	}

	for _, suite := range suites {
		suiteName := suite.Name
		if s.FeatureRegex != nil {
			matches := s.FeatureRegex.FindStringSubmatch(suite.Name)
			if len(matches) < 2 {
				slog.Info("suite did not match feature regex for sdk, skipping", "sdk", s.Name, "suite", suite.Name, "matches", matches)
				continue
			}
			suiteName = matches[1]
			slog.Info("regex success for suite", "sdk", s.Name, "before", suite.Name, "after", suiteName)
		}

		if knownVectors[suiteName] == nil {
			slog.Info("ignoring test suite that does not correspond to known feature", "suite", suiteName)
			continue
		}

		for _, test := range suite.Tests {
			testName := test.Name
			if s.VectorRegex != nil {
				matches := s.VectorRegex.FindStringSubmatch(test.Name)
				if len(matches) < 2 {
					slog.Info("test did not match feature regex for sdk, skipping", "sdk", s.Name, "suite", suiteName, "test", test.Name, "matches", matches)
					continue
				}
				testName = matches[1]
				slog.Info("regex success for test", "sdk", s.Name, "before", test.Name, "after", testName)
			}

			if !knownVectors[suiteName][testName] {
				slog.Info("ignoring test that does not correspond to known vector", "suite", suiteName, "test", testName)
				continue
			}

			errs := []error{}
			if test.Error != nil {
				errs = append(errs, test.Error)
			}

			results[suiteName][testName] = Result{
				Exists: true,
				Errors: errs,
				Time:   test.Duration,
			}
		}
	}

	return Report{
		SDK:     s,
		Results: results,
	}, nil
}
