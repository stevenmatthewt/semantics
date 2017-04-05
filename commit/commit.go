package commit

import (
	"strings"

	"github.com/stevenmatthewt/semantics/bump"
)

type Commits struct {
	Commits []Commit
}

type Commit struct {
	Message string
	Hash    string
}

func (c Commits) ScanForBumps(bumpMap bump.Map) (bumps []bump.Bump) {
	for _, commit := range c.Commits {
		if strings.Index(commit.Message, bumpMap.Major+":") == 0 {
			bumps = append(bumps, bump.MajorBump{})
		}
		if strings.Index(commit.Message, bumpMap.Minor+":") == 0 {
			bumps = append(bumps, bump.MinorBump{})
		}
		if strings.Index(commit.Message, bumpMap.Patch+":") == 0 {
			bumps = append(bumps, bump.PatchBump{})
		}
	}

	return bumps
}
