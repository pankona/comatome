package comatome

import (
	"context"
	"fmt"

	"github.com/google/go-github/v24/github"
)

// QueryCreatedRepos queries created repositories on specified term (fromto)
func QueryCreatedRepos(c *Client, fromto *FromTo) ([]string, error) {
	name := Username(c)
	page := 1
	createdRepos := make([]string, 0)
	for {
		from, to := fromto.QueryStr()
		result, resp, err := c.Search.Repositories(
			context.Background(),
			fmt.Sprintf("user:%s created:%s..%s fork:true", name, from, to),
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

// ShowCreatedRepos shows created repositories
func ShowCreatedRepos(repos []string) {
	fmt.Printf("Created %d repositories\n", len(repos))
	for _, v := range repos {
		fmt.Printf("%s\n", v)
	}
}
