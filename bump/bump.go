package bump

import (
	"fmt"

	"github.com/stevenmatthewt/semantics/tag"
)

type Bump interface {
	Bump(t tag.Tag) tag.Tag
}

type MajorBump struct{}
type MinorBump struct{}
type PatchBump struct{}

func (b MajorBump) Bump(t tag.Tag) tag.Tag {
	fmt.Printf("Bumping tag (major): %+v\n", t)
	t.Major++
	t.Minor = 0
	t.Patch = 0

	return t
}

func (b MinorBump) Bump(t tag.Tag) tag.Tag {
	fmt.Printf("Bumping tag (minor): %+v\n", t)
	t.Minor++
	t.Patch = 0

	return t
}

func (b PatchBump) Bump(t tag.Tag) tag.Tag {
	fmt.Printf("Bumping tag (patch): %+v\n", t)
	t.Patch++

	return t
}
