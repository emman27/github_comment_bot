package actions

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
)

// CreateComment creates a comment with a given text
func CreateComment(repoOwner, repo string, prNumber int, commentText string) error {
	ctx := context.TODO()
	comment := &github.IssueComment{
		Body: &commentText,
	}
	_, _, err := githubClient.Issues.CreateComment(ctx, repoOwner, repo, prNumber, comment)
	if err != nil {
		logrus.Error("Posting comment failed: ", err)
	}
	return err
}
