package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/v21/github"
	"golang.org/x/oauth2"
)

func main() {
	token := os.Getenv("GITHUB_API_TOKEN")

	var tc *http.Client
	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc = oauth2.NewClient(oauth2.NoContext, ts)
	}

	c := github.NewClient(tc)

	page := 1
	for {
		events, resp, err := c.Activity.ListEventsPerformedByUser(
			context.Background(),
			"pankona",
			false, // publicOnly
			&github.ListOptions{Page: page})
		if err != nil {
			panic(err)
		}

		if resp.NextPage == 0 {
			break
		}

		page = resp.NextPage

		for _, v := range events {
			fmt.Printf("%v: %s: %s\n", *v.CreatedAt, *v.Type, *v.Repo.Name)
		}
	}
}
