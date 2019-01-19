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
		result, resp, err := c.Search.Repositories(
			context.Background(),
			"user:pankona+created:2019-01-01..2019-02-01",
			&github.SearchOptions{
				ListOptions: github.ListOptions{
					Page: page,
				}})
		if err != nil {
			panic(err)
		}

		page = resp.NextPage

		for _, v := range result.Repositories {
			fmt.Printf("%s: createdAt: %v\n", *v.Name, *v.CreatedAt)
		}

		if resp.NextPage == 0 {
			break
		}
	}
}
