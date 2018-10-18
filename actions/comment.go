package actions

import (
	"context"

	"github.com/google/go-github/github"
)

// CreateComment creates a comment with a given text
func CreateComment(repoOwner, repo string, prNumber int, commentText string) (*github.IssueComment, *github.Response, error) {
	ctx := context.TODO()
	comment := &github.IssueComment{
		Body: &commentText,
	}
	return githubClient.Issues.CreateComment(ctx, repoOwner, repo, prNumber, comment)
}
