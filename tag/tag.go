package tag

import "fmt"

type Tag struct {
	Major int
	Minor int
	Patch int
}

type Getter interface {
	GetLatestTag() Tag
}

func (t Tag) Tag() string {
	return fmt.Sprintf("%d.%d.%d", t.Major, t.Minor, t.Patch)
}
