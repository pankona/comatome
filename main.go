package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sort"

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

func QueryOpenedPullRequests(c *Client, fromto string) (map[string]int, error) {
	page := 1
	pulls := make(map[string]int)
	for {
		result, resp, err := c.Search.Issues(
			context.Background(),
			fmt.Sprintf("type:pr+author:pankona+created:%s", fromto),
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
			pulls[*v.RepositoryURL] += 1
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
	fmt.Printf("%d pull requests opened\n", total)
}
