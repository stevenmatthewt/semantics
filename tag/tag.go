package tag

type Tag struct {
	Tag string
}

type Getter interface {
	GetLatestTag() Tag
}
