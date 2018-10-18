package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/emman27/github_comment_bot/server"
	"github.com/sirupsen/logrus"
)

var port = flag.Int("port", 8000, "Port to run server on")

func main() {
	http.HandleFunc("/", server.Handler)
	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
