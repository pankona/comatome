package comatome

import (
	"context"
	"net/http"

	"github.com/google/go-github/v24/github"
	"golang.org/x/oauth2"
)

// Client manages communication with GitHub API
type Client struct {
	*github.Client
}

// NewClient returns a HTTP Client with specified GitHub API token
func NewClient(token string) *Client {
	var tc *http.Client
	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)

		tc = oauth2.NewClient(context.Background(), ts)
	}

	return &Client{github.NewClient(tc)}
}
