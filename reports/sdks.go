package reports

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/exp/slog"
)

var (
	SDKs = []SDKMeta{
		{
			Name:         "web5-js",
			Repo:         "TBD54566975/web5-js",
			ArtifactName: "junit-results",
			FeatureRegex: regexp.MustCompile(`Web5TestVectors(\w+)`),
			VectorRegex:  regexp.MustCompile(`.* Web5TestVectors\w+ (\w+)`),
			VectorPath:   "test-vectors",
			Type:         "web5",
		},
		{
			Name:         "web5-kt",
			Repo:         "TBD54566975/web5-kt",
			ArtifactName: "test-results",
			FeatureRegex: regexp.MustCompile(`web5\.sdk\.\w+.Web5TestVectors(\w+)`),
			VectorRegex:  regexp.MustCompile(`(\w+)\(\)`),
			VectorPath:   "test-vectors",
			Type:         "web5",
		},
		{
			Name:         "tbdex-js",
			Repo:         "TBD54566975/tbdex-js",
			ArtifactName: "junit-results",
			FeatureRegex: regexp.MustCompile(`TbdexTestVectors(\w+)`),
			VectorRegex:  regexp.MustCompile(`\w+ \w+ (\w+)`),
			VectorPath:   "test-vectors",
			Type:         "tbdex",
		},
		{
			Name:         "tbdex-kt",
			Repo:         "TBD54566975/tbdex-kt",
			ArtifactName: "test-results",
			FeatureRegex: regexp.MustCompile(`tbdex\.sdk\.\w+.TbdexTestVectors(\w+)`),
			VectorRegex:  regexp.MustCompile(`(\w+)\(\)`),
			VectorPath:   "test-vectors",
			Type:         "tbdex",
		},
	}
)

func GetAllReports() ([]Report, error) {
	ctx := context.Background()

	var reports []Report
	for _, sdk := range SDKs {
		artifact, err := downloadArtifact(ctx, sdk)
		if err != nil {
			return nil, fmt.Errorf("error downloading artifact from %s: %v", sdk.Repo, err)
		}

		suites, err := readArtifactZip(artifact)
		if err != nil {
			return nil, fmt.Errorf("error parsing artifact from %s: %v", sdk.Repo, err)
		}

		report, err := sdk.buildReport(suites)
		if err != nil {
			return nil, fmt.Errorf("error processing data from %s: %v", sdk.Repo, err)
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func downloadArtifact(ctx context.Context, sdk SDKMeta) ([]byte, error) {
	owner, repo, _ := strings.Cut(sdk.Repo, "/")
	artifacts, _, err := gh.Actions.ListArtifacts(ctx, owner, repo, nil)
	if err != nil {
		slog.Error("error listing artifacts", "")
		return nil, fmt.Errorf("error getting artifact list: %v", err)
	}

	var artifactURL string
	for _, a := range artifacts.Artifacts {
		if a.GetWorkflowRun().GetHeadBranch() != "main" {
			continue
		}
		if *a.Name == sdk.ArtifactName {
			artifactURL = *a.ArchiveDownloadURL
			slog.Info("downloading artifact", "repo", sdk.Repo, "commit", a.GetWorkflowRun().GetHeadSHA())
			break
		}
	}

	req, err := http.NewRequest(http.MethodGet, artifactURL, nil)
	if err != nil {
		return nil, err
	}
	bearer := ghToken
	if ghToken == "" {
		bearer, err = ghTransport.Token(ctx)
		if err != nil {
			return nil, fmt.Errorf("error getting github token: %v", err)
		}
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearer))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making http request to %s: %v", artifactURL, err)
	}
	defer resp.Body.Close()

	artifact, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	slog.Info("downloaded artifact", "sdk", sdk.Repo, "size", len(artifact))

	return artifact, nil
}
