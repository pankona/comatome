package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v21/github"
)

func QueryCreatedRepos(c *Client, fromto string) ([]string, error) {
	name := Username(c)
	page := 1
	createdRepos := make([]string, 0)
	for {
		result, resp, err := c.Search.Repositories(
			context.Background(),
			fmt.Sprintf("user:%s created:%s fork:true", name, fromto),
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

func ShowCreatedRepos(repos []string) {
	fmt.Printf("Created %d repositories\n", len(repos))
	for _, v := range repos {
		fmt.Printf("%s\n", v)
	}
}
