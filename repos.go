package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v21/github"
)

func QueryCreatedRepos(c *Client, fromto string) ([]string, error) {
	page := 1
	createdRepos := make([]string, 0)
	for {
		result, resp, err := c.Search.Repositories(
			context.Background(),
			fmt.Sprintf("user:pankona created:%s", fromto),
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
