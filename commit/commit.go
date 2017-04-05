package commit

import (
	"github.com/stevenmatthewt/semantics/bump"
)

// Commits stores a list of commits
type Commits struct {
	Commits []Commit
}

// Commit represents a Git Commit object
type Commit struct {
	Message string
	Hash    string
}

// ScanForBumps scans for any major, minor, or patch updates that should
// occur in the list of commits.
func (c Commits) ScanForBumps(bumpMap bump.Map) (bumps []bump.Bump) {
	for _, commit := range c.Commits {
		if bumpMap.Major.MatchString(commit.Message) {
			bumps = append(bumps, bump.MajorBump{})
		}
		if bumpMap.Minor.MatchString(commit.Message) {
			bumps = append(bumps, bump.MinorBump{})
		}
		if bumpMap.Patch.MatchString(commit.Message) {
			bumps = append(bumps, bump.PatchBump{})
		}
	}

	return bumps
}
