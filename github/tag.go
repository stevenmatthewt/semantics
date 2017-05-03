package github

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/stevenmatthewt/semantics/tag"
)

// GetLatestTag returns the latest tag that exists according to git
func (g GitHub) GetLatestTag() (tag.Tag, error) {

}

// PushTag pushes the specified tag to the remote
func (g GitHub) PushTag(t tag.Tag) (err error) {
	client := github.NewClient(nil)
	repo := &github.Repository{

		Name:    github.String("stevenmatthewt/semantics"),
		Private: github.Bool(true),
	}
	githubTag := github.Tag{
		Tag:    github.String(t.String()),
		Object: github.String(""),
	}
	service := &github.GitService{}
	service.CreateTag(context.Background(), g.owner, g.repo, githubTag)
}
