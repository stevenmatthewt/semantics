package tag

type Tag struct {
	Tag   string
	Major int
	Minor int
	Patch int
}

type Getter interface {
	GetLatestTag() Tag
}
