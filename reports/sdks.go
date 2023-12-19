package reports

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/google/go-github/v57/github"
	"golang.org/x/exp/slog"
)

var (
	ghToken = os.Getenv("GITHUB_TOKEN")
	gh      = github.NewClient(nil).WithAuthToken(ghToken)
	SDKs    = []SDKMeta{
		{
			Name:         "web5-js",
			Repo:         "TBD54566975/web5-js",
			ArtifactName: "junit-results",
			FeatureRegex: regexp.MustCompile(`Web5TestVectors(\w+)`),
			VectorRegex:  regexp.MustCompile(`\w+ \w+ (\w+)`),
		},
		{
			Name:         "web5-kt",
			Repo:         "TBD54566975/web5-kt",
			ArtifactName: "test-results",
			FeatureRegex: regexp.MustCompile(`web5\.sdk\.\w+.Web5TestVectors(\w+)`),
			VectorRegex:  regexp.MustCompile(`(\w+)\(\)`),
		},
	}
)

func init() {
	if ghToken == "" {
		panic("please set environment variable GITHUB_TOKEN to a valid github token (generate one at https://github.com/settings/tokens?type=beta)")
	}
}

func GetAllReports() ([]Report, error) {
	ctx := context.TODO()

	reports := []Report{}
	for _, sdk := range SDKs {
		artifact, err := downloadArtifact(ctx, sdk)
		if err != nil {
			return nil, err
		}

		suites, err := readArtifactZip(artifact)
		if err != nil {
			return nil, err
		}

		report, err := sdk.buildReport(suites)
		if err != nil {
			return nil, err
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
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ghToken))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	artifact, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	slog.Info("downloaded artifact", "sdk", sdk.Repo, "size", len(artifact))

	return artifact, nil
}
