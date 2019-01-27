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
	showCreatedRepos(results)
}

func openedPullRequests(c *Client, fromto string) {
	results, err := QueryOpenedPullRequests(c, fromto)
	if err != nil {
		panic(err)
	}
	ShowOpenedPullRequests(results)
}
