package receivers

import (
	"strings"

	"github.com/emman27/github_comment_bot/actions"
	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
)

// Commands that can be given in Github Comments
const (
	CommandCI        = "/ci"
	MessageCIStarted = "I hear you bro, starting CI now"
)

// ReceiveIssueComment deals with the event of a comment being made on a PR or issue
// It will do the following actions:
func ReceiveIssueComment(e *github.IssueCommentEvent) (map[string]interface{}, error) {
	body := *(e.Comment.Body)
	logrus.Debug(body)
	if strings.HasPrefix(body, CommandCI) {
		logrus.Info("Triggered CI")
		actions.CreateComment(*e.Repo.Owner.Name, *e.Repo.Name, int(*e.Issue.ID), MessageCIStarted)
	}
	return map[string]interface{}{}, nil
}
