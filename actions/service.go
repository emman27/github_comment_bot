package actions

import (
	"net/http"
	"net/url"
	"os"

	"github.com/google/go-github/github"
)

var (
	username = os.Getenv("GITHUB_USERNAME")
	password = os.Getenv("GITHUB_PASSWORD")
)

var githubClient = github.NewClient(&http.Client{
	Transport: &http.Transport{
		Proxy: setBasicAuth,
	},
})

func setBasicAuth(req *http.Request) (*url.URL, error) {
	req.SetBasicAuth(username, password)
	return nil, nil
}
