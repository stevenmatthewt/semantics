package tag

import "fmt"

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
	return fmt.Sprintf("%d.%d.%d", t.Major, t.Minor, t.Patch)
}
