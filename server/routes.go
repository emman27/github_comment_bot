package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/emman27/github_comment_bot/receivers"
	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
)

// Handler is the receiving endpoint for all webhooks
func Handler(w http.ResponseWriter, r *http.Request) {
	// payload, err := github.ValidatePayload(r, []byte{})
	// if err != nil {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	w.Write(ErrorInvalidSignature)
	// 	return
	// }
	var data map[string]interface{}
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrorInvalidPayload)
	}
	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrorInvalidPayload)
		return
	}
	switch event := event.(type) {
	case *github.IssueCommentEvent:
		data, err = receivers.ReceiveIssueComment(event)
	default:
		data, err = map[string]interface{}{
			"error": "Unknown event type",
		}, nil
	}
	body, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ErrorInternalServer)
	}
	fmt.Fprintf(w, string(body))
}
