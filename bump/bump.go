package bump

import (
	"fmt"

	"github.com/stevenmatthewt/semantics/tag"
)

// Bump represents an object that is capable of taking a tag
// and increasing it's version according to Semantic Versioning
type Bump interface {
	Bump(t tag.Tag) tag.Tag
}

// MajorBump increases the major version of a tag
type MajorBump struct{}

// MinorBump increases the minor version of a tag
type MinorBump struct{}

// PatchBump increases the patch version of a tag
type PatchBump struct{}

// Bump takes a tag and increases it by on major version
func (b MajorBump) Bump(t tag.Tag) tag.Tag {
	fmt.Printf("Bumping tag (major): %+v\n", t)
	t.Major++
	t.Minor = 0
	t.Patch = 0

	return t
}

// Bump takes a tag and increases it by on minor version
func (b MinorBump) Bump(t tag.Tag) tag.Tag {
	fmt.Printf("Bumping tag (minor): %+v\n", t)
	t.Minor++
	t.Patch = 0

	return t
}

// Bump takes a tag and increases it by on patch version
func (b PatchBump) Bump(t tag.Tag) tag.Tag {
	fmt.Printf("Bumping tag (patch): %+v\n", t)
	t.Patch++

	return t
}
