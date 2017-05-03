package bump

import (
	"reflect"
	"testing"

	"github.com/stevenmatthewt/semantics/tag"
)

type test struct {
	name     string
	startTag tag.Tag
	bumps    []Bump
	result   tag.Tag
}

func TestBumping(t *testing.T) {
	table := []test{
		test{
			"test single bump",
			tag.Tag{
				Major: 1,
				Minor: 2,
				Patch: 3,
			},
			[]Bump{
				PatchBump{},
			},
			tag.Tag{
				Major: 1,
				Minor: 2,
				Patch: 4,
			},
		},
		test{
			"test multi bump",
			tag.Tag{
				Major: 1,
				Minor: 2,
				Patch: 3,
			},
			[]Bump{
				PatchBump{},
				MinorBump{},
				MajorBump{},
				MinorBump{},
				MinorBump{},
				PatchBump{},
				MinorBump{},
				PatchBump{},
				PatchBump{},
			},
			tag.Tag{
				Major: 2,
				Minor: 3,
				Patch: 2,
			},
		},
		test{
			"test no bump",
			tag.Tag{
				Major: 1,
				Minor: 2,
				Patch: 3,
			},
			[]Bump{},
			tag.Tag{
				Major: 1,
				Minor: 2,
				Patch: 3,
			},
		},
	}

	for _, test := range table {
		tag := test.startTag
		for _, b := range test.bumps {
			tag = b.Bump(tag)
		}
		if !reflect.DeepEqual(tag, test.result) {
			t.Errorf("%s: got %+v want %+v", test.name, tag, test.result)
		}
	}
}
