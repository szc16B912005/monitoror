package gogithub

import (
	"context"

	githubApi "github.com/google/go-github/github"
)

type PullRequestService interface {
	List(ctx context.Context, owner string, repo string, opt *githubApi.PullRequestListOptions) ([]*githubApi.PullRequest, *githubApi.Response, error)
}
