package commit

import (
	"reflect"
	"testing"

	"github.com/stevenmatthewt/semantics/bump"
)

type test struct {
	name    string
	commits Commits
	bumpMap bump.Map
	result  []bump.Bump
}

func TestScanForBumps(t *testing.T) {
	defaultBumpMap, err := bump.MapFromStrings("^major:.*", "^minor:.*", "^patch:.*")
	if err != nil {
		t.Errorf("Failed to create bumpMap: %v", err)
	}
	table := []test{
		test{
			"Test single patch",
			Commits{
				[]Commit{
					Commit{
						Message: "patch: hello",
					},
				},
			},
			defaultBumpMap,
			[]bump.Bump{
				bump.PatchBump{},
			},
		},
		test{
			"Test multiple patch",
			Commits{
				[]Commit{
					Commit{
						Message: "patch: hello",
					},
					Commit{
						Message: "patch: hi there",
					},
					Commit{
						Message: "patch: howdy",
					},
					Commit{
						Message: "patch: Ciao",
					},
				},
			},
			defaultBumpMap,
			[]bump.Bump{
				bump.PatchBump{},
				bump.PatchBump{},
				bump.PatchBump{},
				bump.PatchBump{},
			},
		},
		test{
			"Test multiple patch",
			Commits{
				[]Commit{
					Commit{
						Message: "patch: hello",
					},
					Commit{
						Message: "minor: hi there",
					},
					Commit{
						Message: "patch: howdy",
					},
					Commit{
						Message: "major: Ciao",
					},
				},
			},
			defaultBumpMap,
			[]bump.Bump{
				bump.PatchBump{},
				bump.MinorBump{},
				bump.PatchBump{},
				bump.MajorBump{},
			},
		},
		test{
			"Test tricky :)",
			Commits{
				[]Commit{
					Commit{
						Message: "patch hello",
					},
					Commit{
						Message: "mInor: hi there",
					},
					Commit{
						Message: "patsh: howdy",
					},
					Commit{
						Message: "major: Ciao",
					},
				},
			},
			defaultBumpMap,
			[]bump.Bump{
				bump.MajorBump{},
			},
		},
	}

	for _, test := range table {
		actual := test.commits.ScanForBumps(test.bumpMap)
		if !reflect.DeepEqual(actual, test.result) {
			t.Errorf("%s: got %+v want %+v", test.name, actual, test.result)
		}
	}
}
