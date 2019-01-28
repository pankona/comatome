package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fromto := flag.String("fromto", "", "specify year and month by yyyymm format to fetch contributions")
	flag.Parse()

	if *fromto == "" {
		panic("fromto is not specified.")
	}
	c := NewClient(os.Getenv("GITHUB_API_TOKEN"))

	commits(c, *fromto)
	fmt.Println()
	createdRepos(c, *fromto)
	fmt.Println()
	openedPullRequests(c, *fromto)
	fmt.Println()
	reviewedPullRequests(c, *fromto)
	fmt.Println()
	openedIssues(c, *fromto)
}

func commits(c *Client, fromto string) {
	results, err := QueryCommitsPerRepo(c, fromto)
	if err != nil {
		panic(err)
	}
	ShowCommitsPerRepo(results)
}

func createdRepos(c *Client, fromto string) {
	results, err := QueryCreatedRepos(c, fromto)
	if err != nil {
		panic(err)
	}
	ShowCreatedRepos(results)
}

func openedPullRequests(c *Client, fromto string) {
	results, err := QueryOpenedPullRequests(c, fromto)
	if err != nil {
		panic(err)
	}
	ShowOpenedPullRequests(results)
}

func reviewedPullRequests(c *Client, fromto string) {
	results, err := QueryReviewedPullRequests(c, fromto)
	if err != nil {
		panic(err)
	}
	ShowReviewedPullRequests(results)
}

func openedIssues(c *Client, fromto string) {
	results, err := QueryOpenedIssues(c, fromto)
	if err != nil {
		panic(err)
	}
	ShowOpenedIssues(results)
}
