package bump

import (
	"github.com/cbdr/semantics/tag"
)

type Bump interface {
	Bump(t tag.Tag) tag.Tag
}

type MajorBump struct{}
type MinorBump struct{}
type PatchBump struct{}

func (b MajorBump) Bump(t tag.Tag) tag.Tag {
	t.Major++
	t.Minor = 0
	t.Patch = 0

	return t
}

func (b MinorBump) Bump(t tag.Tag) tag.Tag {
	t.Minor++
	t.Patch = 0

	return t
}

func (b PatchBump) Bump(t tag.Tag) tag.Tag {
	t.Patch++

	return t
}
