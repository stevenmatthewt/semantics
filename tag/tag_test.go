package tag

import (
	"reflect"
	"testing"
)

type fromStringTest struct {
	name      string
	tag       string
	result    Tag
	expectErr bool
}

func TestFromString(t *testing.T) {
	table := []fromStringTest{
		fromStringTest{
			"Test 1",
			"v0.0.0",
			Tag{0, 0, 0},
			false,
		},
		fromStringTest{
			"Test 2",
			"v1.0.4",
			Tag{1, 0, 4},
			false,
		},
		fromStringTest{
			"Test 3",
			"0.0.0",
			Tag{},
			true,
		},
		fromStringTest{
			"Test 4",
			"v1.3.a",
			Tag{},
			true,
		},
		fromStringTest{
			"Test 5",
			"v0.0",
			Tag{},
			true,
		},
	}

	for _, test := range table {
		expected, err := FromString(test.tag)
		if test.expectErr == true && err == nil {
			t.Error("Expected error but did not get one")
		} else if test.expectErr == false && err != nil {
			t.Errorf("Got an unexpected error: %v", err)
		}
		if !reflect.DeepEqual(expected, test.result) {
			t.Errorf("%s: got %+v want %+v", test.name, expected, test.result)
		}
		if test.expectErr {
			continue
		}
		if test.result.String() != test.tag {
			t.Errorf("%s: got %+v want %+v", test.name, test.tag, test.result.String())
		}
	}
}
