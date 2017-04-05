package commit

import (
	"strings"

	"github.com/cbdr/semantics/bump"
)

type Commits struct {
	Commits []Commit
}

type Commit struct {
	Message string
	Hash    string
}

func (c Commits) ScanForBumps(bumpMap bump.Map) (major int, minor int, patch int) {
	for _, commit := range c.Commits {
		if strings.Index(commit.Message, bumpMap.Major+":") == 0 {
			major++
		}
		if strings.Index(commit.Message, bumpMap.Minor+":") == 0 {
			minor++
		}
		if strings.Index(commit.Message, bumpMap.Patch+":") == 0 {
			patch++
		}
	}

	return major, minor, patch
}
