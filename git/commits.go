package git

import (
	"log"
	"os/exec"
	"strings"

	"github.com/cbdr/semantics/tag"
)

func GetCommitsSinceTag(t tag.Tag) []string {
	//git log --pretty=oneline head...tag
	commits, err := runGitLog(t.Tag)
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
