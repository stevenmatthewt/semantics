package git

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/stevenmatthewt/semantics/output"
	"github.com/stevenmatthewt/semantics/tag"
)

const invalidTagFormat = "Tag %s is not a valid format"

// GetLatestTag returns the latest tag that exists according to git
func (g Git) GetLatestTag() (tag.Tag, error) {
	latestTag, err := runGitDescribe()
	if err != nil {
		output.Fatal(fmt.Errorf("Failed to fetch tag: %v", err))
	}

	return tagStringToTag(latestTag)
}

// PushTag pushes the specified tag to the remote
func (g Git) PushTag(t tag.Tag) (err error) {
	cmd := exec.Command("git", "tag", t.String())
	_, err = runCommand(cmd)
	if err != nil {
		return err
	}

	cmd = exec.Command("git", "push", "origin", t.String())
	_, err = runCommand(cmd)
	if err != nil {
		return err
	}
	return nil
}

func runGitDescribe() (string, error) {
	// The following glob(7) pattern is not perfect. It will match things like v1.4badstring.8
	// But it narrows down the results by a good bit
	cmd := exec.Command("git", "describe", "--abbrev=0", "--tags", "--match=v[0-9]*\\.[0-9]*\\.[0-9]*")
	// TODO: if we find a tag, but it's invalid (human created), we should retry and find the one previous.
	// Probably with a `git describe <bad tag that was found>`
	return runCommand(cmd)
}

func tagStringToTag(tagString string) (tag.Tag, error) {
	if len(tagString) > 0 {
		tagString = tagString[1:]
	}
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
