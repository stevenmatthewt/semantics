package git

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/stevenmatthewt/semantics/tag"
)

const invalidTagFormat = "Tag %s is not a valid format"

func GetLatestTag() (tag.Tag, error) {
	latestTag, err := runGitDescribe()
	if err != nil {
		log.Fatal(err)
	}

	return tagStringToTag(latestTag)
}

func PushTag(t tag.Tag) (err error) {
	cmd := exec.Command("git", "tag", t.Tag())
	_, err = runCommand(cmd)
	if err != nil {
		return err
	}

	cmd = exec.Command("git", "push", "origin", t.Tag())
	_, err = runCommand(cmd)
	if err != nil {
		return err
	}
	return nil
}

func runGitDescribe() (string, error) {
	cmd := exec.Command("git", "describe", "--abbrev=0", "--tags")
	return runCommand(cmd)
}

func tagStringToTag(tagString string) (tag.Tag, error) {
	tagArray := strings.Split(tagString, ".")
	if len(tagArray) != 3 {
		return tag.Tag{}, errors.New("Latest fetched tag was not in proper a.b.c format")
	}
	major, err := strconv.Atoi(tagArray[0])
	if err != nil {
		return tag.Tag{}, fmt.Errorf(invalidTagFormat, tagArray[0])
	}
	minor, err := strconv.Atoi(tagArray[1])
	if err != nil {
		return tag.Tag{}, fmt.Errorf(invalidTagFormat, tagArray[1])
	}
	patch, err := strconv.Atoi(tagArray[2])
	if err != nil {
		return tag.Tag{}, fmt.Errorf(invalidTagFormat, tagArray[2])
	}
	return tag.Tag{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}
