package main

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/google/go-github/v21/github"
)

func QueryOpenedIssues(c *Client, fromto string) (map[string]int, error) {
	name := Username(c)
	page := 1
	pulls := make(map[string]int)
	for {
		result, resp, err := c.Search.Issues(
			context.Background(),
			fmt.Sprintf("type:issue author:%s created:%s", name, fromto),
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
			pulls[repo]++
		}

		if resp.NextPage == 0 {
			break
		}
	}
	return pulls, nil
}

func ShowOpenedIssues(pulls map[string]int) {
	keys := make([]string, len(pulls))
	index := 0
	for k := range pulls {
		keys[index] = k
		index++
	}
	sort.Strings(keys)

	total := 0
	for _, v := range keys {
		total += pulls[v]
	}
	fmt.Printf("Opened %d issues in %d repositories\n", total, len(pulls))
	for _, v := range keys {
		fmt.Printf("%d\t%s\n", pulls[v], v)
	}
}
