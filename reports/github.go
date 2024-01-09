package reports

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	ghinstallation "github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v57/github"
	"golang.org/x/exp/slog"
)

var (
	ghAppName                 = os.Getenv("CICD_ROBOT_GITHUB_APP_NAME")
	ghAppPrivateKey           = os.Getenv("CICD_ROBOT_GITHUB_APP_PRIVATE_KEY")
	ghAppIDString             = os.Getenv("CICD_ROBOT_GITHUB_APP_ID")
	ghAppInstallationIDString = os.Getenv("CICD_ROBOT_GITHUB_APP_INSTALLATION_ID")

	ghToken = os.Getenv("GITHUB_TOKEN")

	gh          *github.Client
	ghTransport *ghinstallation.Transport

	ghUserName = fmt.Sprintf("%s[bot]", ghAppName)
)

func init() {
	if ghToken != "" {
		slog.Info("using GITHUB_TOKEN for auth")
		gh = github.NewTokenClient(context.Background(), ghToken)
		return
	}

	ghAppID, err := strconv.ParseInt(ghAppIDString, 10, 32)
	if err != nil {
		slog.Error("invalid or unset app ID. Please set environment variable CICD_ROBOT_GITHUB_APP_ID to a valid integer")
		panic(err)
	}

	ghInstallationID, err := strconv.ParseInt(ghAppInstallationIDString, 10, 32)
	if err != nil {
		slog.Error("invalid or unset installation ID. Please set environment variable CICD_ROBOT_GITHUB_APP_INSTALLATION_ID to a valid integer")
		panic(err)
	}

	ghTransport, err = ghinstallation.New(http.DefaultTransport, ghAppID, ghInstallationID, []byte(ghAppPrivateKey))
	if err != nil {
		slog.Error("error initializing github auth transport.")
		panic(err)
	}

	gh = github.NewClient(&http.Client{Transport: ghTransport})

	user, _, err := gh.Users.Get(context.Background(), ghUserName)
	if err != nil {
		slog.Error("error getting own (app) user info")
		panic(err)
	}

	gitConfig["user.email"] = fmt.Sprintf("%d+%s@users.noreply.github.com", user.GetID(), ghUserName)
	gitConfig["user.name"] = ghUserName
}
