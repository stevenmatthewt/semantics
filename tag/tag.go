package tag

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const invalidTagFormat = "Tag section %s is not a valid format"

// Tag represents a Semantic Release Tag
type Tag struct {
	Major int
	Minor int
	Patch int
}

// Getter is an interface around anything that is capable of retreiving tag data.
type Getter interface {
	GetLatestTag() Tag
}

// String returns a string representation of a Tag
func (t Tag) String() string {
	return fmt.Sprintf("v%d.%d.%d", t.Major, t.Minor, t.Patch)
}

// FromString takes a string and converts it to a Tag, if possible
func FromString(tagString string) (Tag, error) {
	if len(tagString) > 0 {
		tagString = tagString[1:]
	}
	tagArray := strings.Split(tagString, ".")
	if len(tagArray) != 3 {
		return Tag{}, errors.New("Tag is not in proper a.b.c format")
	}
	major, err := strconv.Atoi(tagArray[0])
	if err != nil {
		return Tag{}, fmt.Errorf(invalidTagFormat, tagArray[0])
	}
	minor, err := strconv.Atoi(tagArray[1])
	if err != nil {
		return Tag{}, fmt.Errorf(invalidTagFormat, tagArray[1])
	}
	patch, err := strconv.Atoi(tagArray[2])
	if err != nil {
		return Tag{}, fmt.Errorf(invalidTagFormat, tagArray[2])
	}
	return Tag{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}
