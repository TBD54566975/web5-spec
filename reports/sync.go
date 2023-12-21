package reports

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v57/github"
	"golang.org/x/exp/slog"
)

const (
	vectorUpdateBranch        = "vector-update"
	gitConfigCredentialHelper = "credential.helper"
	vectorUpdateCommitMessage = "update test vectors"
)

// these should be consts but the library expects a pointer to a string, which cannot be done with a const
var (
	vectorUpdatePRTitle      = "Update Test Vectors - Out of Sync"
	vectorUpdatePRBody       = "Some test vectors were changed or added to the sdk-development repo, so they need to be updated in this repo. This is an automated PR that keeps the test vectors in sync from the [main test-vectors location](https://github.com/TBD54566975/sdk-development/tree/main/web5-test-vectors)"
	vectorUpdatePRBaseBranch = "main"
)

var gitCredentialStoreFile string

var gitConfig = make(map[string]string)

func SyncSDK(sdk SDKMeta) error {
	slog.Info("syncing vectors", "repo", sdk.Repo)

	tmpdir, err := os.MkdirTemp("", "vector-update")
	if err != nil {
		return fmt.Errorf("error making a temp dir: %v", err)
	}
	defer os.RemoveAll(tmpdir)

	// clone sdk.Repo
	// check if a vector update branch already exists.
	// If vector update branch exists, check it out + rebase on default branch
	// if vector update branch does not exist, make it
	err = clone(fmt.Sprintf("https://github.com/%s", sdk.Repo), tmpdir)
	if err != nil {
		return fmt.Errorf("error cloning repo %s: %v", sdk.Repo, err)
	}

	// copy ../../web5-test-vectors/* to sdk.VectorPath
	err = copyDir("../web5-test-vectors", filepath.Join(tmpdir, sdk.VectorPath))
	if err != nil {
		return fmt.Errorf("error copying current vectors to cloned repo: %v", err)
	}

	// check if git says the repo has changed - return if it hasn't
	err = git("-C", tmpdir, "diff-index", "--quiet", "HEAD")
	if err != nil {
		exitError := new(exec.ExitError)
		if !errors.As(err, &exitError) {
			return fmt.Errorf("error checking if repo changed: %v", err)
		}

		slog.Info("repo changed after copying current vectors in")
	} else {
		slog.Info("repo did not change after copying current vectors in, not taking further action")
		return nil
	}

	// commit
	if err := git("-C", tmpdir, "commit", "-a", "-m", vectorUpdateCommitMessage); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}

	// push
	if err := git("-C", tmpdir, "push", "origin", vectorUpdateBranch, "--force"); err != nil {
		return fmt.Errorf("error pushing changes: %v", err)
	}

	// open a pull request if one isn't already open
	if err := openPRIfNeeded(sdk.Repo); err != nil {
		return fmt.Errorf("error opening PR: %v", err)
	}
	return nil
}

// clone the repo and checkout the correct branch and rebase it on main
func clone(url string, dest string) error {
	if err := git("clone", url, dest); err != nil {
		return err
	}

	if err := git("-C", dest, "checkout", vectorUpdateBranch); err != nil {
		exitError := &exec.ExitError{}
		if !errors.As(err, &exitError) {
			return err
		}

		err = git("-C", dest, "checkout", "-b", vectorUpdateBranch)
		if err != nil {
			return err
		}
	}

	if err := git("-C", dest, "rebase", "main"); err != nil {
		slog.Warn("rebase failed")
		return err
	}

	return nil
}

func git(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Env = []string{fmt.Sprintf("GIT_CONFIG_COUNT=%d", len(gitConfig))}
	i := 0
	for k, v := range gitConfig {
		cmd.Env = append(cmd.Env,
			fmt.Sprintf("GIT_CONFIG_KEY_%d=%s", i, k),
			fmt.Sprintf("GIT_CONFIG_VALUE_%d=%s", i, v),
		)
		i = i + 1
	}

	slog.Info("invoking", "git", args)
	if err := cmd.Run(); err != nil {
		exitError := &exec.ExitError{}
		if !errors.As(err, &exitError) {
			return err
		}
		slog.Info("command did not succeed", "exit_code", exitError.ExitCode())
		return err
	}

	return nil
}

func copyDir(src, dest string) error {
	return filepath.WalkDir(src, func(path string, _ fs.DirEntry, err error) error {
		if err != nil {
			slog.Error("error from walkdir")
			return err
		}

		if !strings.HasSuffix(path, ".json") {
			return nil
		}

		relativePath, _ := filepath.Rel(src, path)
		destPath := filepath.Join(dest, relativePath)

		slog.Info("mkdir", "dir", filepath.Dir(destPath))
		err = os.MkdirAll(filepath.Dir(destPath), 0755)
		if err != nil && !errors.Is(err, os.ErrExist) {
			slog.Error("error creating dir", "dir", destPath)
			return err
		}

		s, err := os.Open(path)
		if err != nil {
			slog.Error("error opening source vector", "file", path)
			return err
		}
		defer s.Close()

		d, err := os.Create(destPath)
		if err != nil {
			slog.Error("error opening dest vector", "file", destPath)
			return err
		}
		defer d.Close()

		_, err = io.Copy(d, s)
		if err != nil {
			slog.Error("error copying vector contents", "src", path, "dest", destPath)
			return err
		}

		if err := git("-C", dest, "add", destPath); err != nil {
			slog.Error("error git add'ing vector", "file", destPath)
			return err
		}

		slog.Info("copied vector", "file", relativePath)
		return nil
	})
}

func ConfigureGitAuth() error {
	slog.Info("telling git about our github token")

	f, err := os.CreateTemp("", "git-credentials")
	if err != nil {
		return err
	}
	gitCredentialStoreFile = f.Name()

	authToken, err := ghTransport.Token(context.Background())
	if err != nil {
		slog.Error("error getting github auth token")
		return err
	}

	cmd := exec.Command("git", "credential-store", "--file", gitCredentialStoreFile, "store")
	cmd.Stdin = strings.NewReader(fmt.Sprintf("protocol=https\nhost=github.com\nusername=%s\npassword=%s", fmt.Sprintf("%s[bot]", ghAppName), authToken))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	gitConfig[gitConfigCredentialHelper] = fmt.Sprintf("store --file %s", gitCredentialStoreFile)

	return nil
}

func CleanupGitAuth() error {
	return os.Remove(gitCredentialStoreFile)
}

func openPRIfNeeded(repo string) error {
	ctx := context.Background()
	owner, repo, _ := strings.Cut(repo, "/")
	head := fmt.Sprintf("%s:%s", owner, vectorUpdateBranch)
	existing, _, err := gh.PullRequests.List(ctx, owner, repo, &github.PullRequestListOptions{
		State: "open",
		Head:  head,
	})
	if err != nil {
		slog.Error("error checking for existing PR")
		return err
	}

	if len(existing) > 0 {
		slog.Info("a PR for that branch already exists, not opening a new one", "pr", existing[0].GetURL())
		return nil
	}

	pr, _, err := gh.PullRequests.Create(ctx, owner, repo, &github.NewPullRequest{
		Title: &vectorUpdatePRTitle,
		Body:  &vectorUpdatePRBody,
		Head:  &head,
		Base:  &vectorUpdatePRBaseBranch,
	})
	if err != nil {
		slog.Error("error creating PR")
		return err
	}

	slog.Info("opened PR", "pr", pr.GetURL())

	return nil
}
