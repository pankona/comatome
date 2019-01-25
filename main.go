package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/go-github/v21/github"
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
}

func commits(c *Client, fromto string) {
	m, err := QueryCommitsPerRepo(c, fromto)
	if err != nil {
		panic(err)
	}
	ShowCommitsPerRepo(m)
}

func createdRepos(c *Client, fromto string) {
	r, err := QueryCreatedRepos(c, fromto)
	if err != nil {
		panic(err)
	}
	showCreatedRepos(r)
}

func QueryCreatedRepos(c *Client, fromto string) ([]string, error) {
	page := 1
	createdRepos := make([]string, 0)
	for {
		result, resp, err := c.Search.Repositories(
			context.Background(),
			fmt.Sprintf("user:pankona+created:%s", fromto),
			&github.SearchOptions{
				ListOptions: github.ListOptions{
					PerPage: 100,
					Page:    page,
				}})
		if err != nil {
			panic(err)
		}

		page = resp.NextPage

		for _, v := range result.Repositories {
			createdRepos = append(createdRepos, *v.FullName)
		}

		if resp.NextPage == 0 {
			break
		}
	}
	return createdRepos, nil
}

func showCreatedRepos(repos []string) {
	for _, v := range repos {
		fmt.Printf("%s\n", v)
	}
	fmt.Printf("%d repositories created\n", len(repos))
}
