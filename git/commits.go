package git

import (
	"errors"
	"log"
	"os/exec"
	"strings"

	"github.com/stevenmatthewt/semantics/commit"
	"github.com/stevenmatthewt/semantics/tag"
)

func GetCommitsSinceTag(t tag.Tag) commit.Commits {
	commitArray, err := runGitLog(t.Tag())
	if err != nil {
		log.Fatal(err)
	}

	commits, err := parseCommitArray(commitArray)
	if err != nil {
		log.Fatal(err)
	}
	return commits
}

func runGitLog(tag string) ([]string, error) {
	cmd := exec.Command("git", "log", "--pretty=oneline", "head..."+tag)
	out, err := runCommand(cmd)
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
