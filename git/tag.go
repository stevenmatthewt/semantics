package git

import (
	"fmt"
	"os/exec"

	"github.com/stevenmatthewt/semantics/output"
	"github.com/stevenmatthewt/semantics/tag"
)

// GetLatestTag returns the latest tag that exists according to git
func (g Git) GetLatestTag() (tag.Tag, error) {
	latestTag, err := runGitDescribe()
	if err != nil {
		output.Fatal(fmt.Errorf("Failed to fetch tag: %v", err))
	}

	return tag.FromString(latestTag)
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
	// But it narrows down the results by a good bit. It will exclude prerelease tags such as `v1.4.8-rc-1`
	cmd := exec.Command("git", "describe", "--abbrev=0", "--tags", "--match=v[0-9]*\\.[0-9]*\\.[0-9]*", "--exclude=v[0-9] *\\.[0 - 9]*\\.[0 - 9]*-*")
	// TODO: if we find a tag, but it's invalid (human created), we should retry and find the one previous.
	// Probably with a `git describe <bad tag that was found>`
	return runCommand(cmd)
}
