package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v21/github"
	"golang.org/x/oauth2"
)

func main() {
	token := os.Getenv("GITHUB_API_TOKEN")
	if token == "" {
		fmt.Println("Please specify GitHub API Token to GITHUB_API_TOKEN environment variable.")
		os.Exit(1)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	c := github.NewClient(tc)
	_, resp, err := c.Activity.ListEventsPerformedByUser(
		context.Background(), "pankona", true, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp: %v\n", resp)
	//fmt.Printf("events: %v\n", events)
}
