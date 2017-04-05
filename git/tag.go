package git

import (
	"log"
	"os/exec"

	"github.com/cbdr/semantics/tag"
)

func GetLatestTag() tag.Tag {
	latestTag, err := runGitDescribe()
	if err != nil {
		log.Fatal(err)
	}
	return tag.Tag{
		Tag: latestTag,
	}
}

func runGitDescribe() (string, error) {
	cmd := exec.Command("git", "describe", "--abbrev=0", "--tags")
	return runCommand(cmd)
}
