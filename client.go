package main

import (
	"net/http"

	"github.com/google/go-github/v21/github"
	"golang.org/x/oauth2"
)

type Client struct {
	*github.Client
}

func NewClient(token string) *Client {
	var tc *http.Client
	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc = oauth2.NewClient(oauth2.NoContext, ts)
	}

	return &Client{github.NewClient(tc)}
}
