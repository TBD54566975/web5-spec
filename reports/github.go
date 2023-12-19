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
	gh              *github.Client
	ghTransport     *ghinstallation.Transport
	ghAppName       = os.Getenv("CICD_ROBOT_GITHUB_APP_NAME")
	ghUserName      = fmt.Sprintf("%s[bot]", ghAppName)
	ghAppPrivateKey = os.Getenv("CICD_ROBOT_GITHUB_APP_PRIVATE_KEY")
)

func init() {
	ghAppID, err := strconv.ParseInt(os.Getenv("CICD_ROBOT_GITHUB_APP_ID"), 10, 32)
	if err != nil {
		slog.Error("invalid or unset app ID. Please set environment variable CICD_ROBOT_GITHUB_APP_ID to a valid integer")
		panic(err)
	}

	ghInstallationID, err := strconv.ParseInt(os.Getenv("CICD_ROBOT_GITHUB_APP_INSTALLATION_ID"), 10, 32)
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
