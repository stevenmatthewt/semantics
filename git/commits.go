package git

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/stevenmatthewt/semantics/commit"
	"github.com/stevenmatthewt/semantics/output"
	"github.com/stevenmatthewt/semantics/tag"
)

// GetCommitsSinceTag returns a Commits object with all commits
// since the provided tag
func GetCommitsSinceTag(t tag.Tag) commit.Commits {
	commitArray, err := runGitLog(t.String())
	if err != nil {
		output.Fatal(err)
	}

	if len(commitArray) == 0 {
		return commit.Commits{}
	}

	commits, err := parseCommitArray(commitArray)
	if err != nil {
		output.Fatal(err)
	}
	return commits
}

func runGitLog(tag string) ([]string, error) {
	cmd := exec.Command("git", "log", "--pretty=oneline", "HEAD..."+tag)
	out, err := runCommand(cmd)
	if out == "" {
		return []string{}, nil
	}
	return strings.Split(out, "\n"), err
}

func parseCommitArray(array []string) (commit.Commits, error) {
	commits := commit.Commits{}
	commits.Commits = make([]commit.Commit, len(array))
	for i, line := range array {
		com, err := parseCommitLine(line)
		if err != nil {
			return commit.Commits{}, err
		}
		// Subtract i to ensure the commits are in proper order
		commits.Commits[len(commits.Commits)-i-1] = com
	}

	return commits, nil
}

func parseCommitLine(line string) (commit.Commit, error) {
	split := strings.SplitN(line, " ", 2)
	if len(split) != 2 {
		return commit.Commit{}, errors.New("Unable to parse commit messages")
	}
	return commit.Commit{
		Hash:    split[0],
		Message: split[1],
	}, nil
}
