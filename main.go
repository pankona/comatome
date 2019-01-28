package main

import (
	"flag"
	"fmt"
	"os"
)

type flags struct {
	commits      *bool
	createdRepos *bool
	openedPRs    *bool
	reviewedPRs  *bool
	openedIssues *bool
}

func main() {
	fromto := flag.String("fromto", "", "specify year and month by yyyymm format to fetch contributions")
	f := flags{
		commits:      flag.Bool("co", false, "show commits"),
		createdRepos: flag.Bool("re", false, "show created repositories"),
		openedPRs:    flag.Bool("op", false, "show opened pull requests"),
		reviewedPRs:  flag.Bool("rp", false, "show reviewed pull requests"),
		openedIssues: flag.Bool("oi", false, "show opened issues"),
	}

	flag.Parse()

	if *fromto == "" {
		panic("fromto is not specified.")
	}
	c := NewClient(os.Getenv("GITHUB_API_TOKEN"))

	if *f.commits {
		commits(c, *fromto)
		fmt.Println()
	}
	if *f.createdRepos {
		createdRepos(c, *fromto)
		fmt.Println()
	}
	if *f.openedPRs {
		openedPullRequests(c, *fromto)
		fmt.Println()
	}
	if *f.reviewedPRs {
		reviewedPullRequests(c, *fromto)
		fmt.Println()
	}
	if *f.openedIssues {
		openedIssues(c, *fromto)
		fmt.Println()
	}
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
