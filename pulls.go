package main

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/google/go-github/v21/github"
)

func QueryOpenedPullRequests(c *Client, fromto string) (map[string]int, error) {
	page := 1
	pulls := make(map[string]int)
	for {
		result, resp, err := c.Search.Issues(
			context.Background(),
			fmt.Sprintf("type:pr author:pankona created:%s", fromto),
			&github.SearchOptions{
				ListOptions: github.ListOptions{
					PerPage: 100,
					Page:    page,
				}})
		if err != nil {
			panic(err)
		}

		page = resp.NextPage

		for _, v := range result.Issues {
			ss := strings.Split(*v.RepositoryURL, "/")
			repo := strings.Join(ss[len(ss)-2:], "/")
			pulls[repo] += 1
		}

		if resp.NextPage == 0 {
			break
		}
	}
	return pulls, nil
}

func ShowOpenedPullRequests(pulls map[string]int) {
	keys := make([]string, len(pulls))
	index := 0
	for k, _ := range pulls {
		keys[index] = k
		index++
	}
	sort.Strings(keys)

	total := 0
	for _, v := range keys {
		fmt.Printf("%d\t%s\n", pulls[v], v)
		total += pulls[v]
	}
	fmt.Printf("%d pull requests opened in %d repositories\n", total, len(pulls))
}

func QueryReviewedPullRequests(c *Client, fromto string) (map[string]int, error) {
	page := 1
	pulls := make(map[string]int)
	for {
		result, resp, err := c.Search.Issues(
			context.Background(),
			fmt.Sprintf("type:pr reviewed-by:pankona updated:%s -author:pankona", fromto),
			&github.SearchOptions{
				ListOptions: github.ListOptions{
					PerPage: 100,
					Page:    page,
				}})
		if err != nil {
			panic(err)
		}

		page = resp.NextPage

		for _, v := range result.Issues {
			ss := strings.Split(*v.RepositoryURL, "/")
			repo := strings.Join(ss[len(ss)-2:], "/")
			pulls[repo] += 1
		}

		if resp.NextPage == 0 {
			break
		}
	}
	return pulls, nil
}

func ShowReviewedPullRequests(pulls map[string]int) {
	keys := make([]string, len(pulls))
	index := 0
	for k, _ := range pulls {
		keys[index] = k
		index++
	}
	sort.Strings(keys)

	total := 0
	for _, v := range keys {
		fmt.Printf("%d\t%s\n", pulls[v], v)
		total += pulls[v]
	}
	fmt.Printf("Reviewed %d pull requests in %d repositories\n", total, len(pulls))
}
