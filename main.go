package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sort"

	"github.com/google/go-github/v21/github"
	"golang.org/x/oauth2"
)

var errIncompleteResult = errors.New("incomplete result error")

func main() {
	token := os.Getenv("GITHUB_API_TOKEN")

	var tc *http.Client
	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc = oauth2.NewClient(oauth2.NoContext, ts)
	}

	c := github.NewClient(tc)

	var m map[string]int
	var err error
	maxRetry := 5

	for i := 0; i < maxRetry; i++ {
		m, err = queryCommitsPerRepo(c)
		if err == errIncompleteResult {
			continue
		}
		break
	}
	if err != nil {
		panic(err)
	}

	keys := make([]string, len(m))
	index := 0
	for k, _ := range m {
		keys[index] = k
		index++
	}
	sort.Strings(keys)

	total := 0
	fmt.Println("commits on 2019-01-01..2019-02-01")
	for _, v := range keys {
		fmt.Printf("%3d commits on %s\n", m[v], v)
		total += m[v]
	}
	fmt.Printf("total %d commits on this month\n", total)
}

func queryCommitsPerRepo(c *github.Client) (map[string]int, error) {
	emails, _, err := c.Users.ListEmails(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	m := make(map[string]int)

	for _, email := range emails {
		page := 1
		for {
			result, resp, err := c.Search.Commits(
				context.Background(),
				fmt.Sprintf("author-email:%s+author-date:2019-01-01..2019-02-01", email.GetEmail()),
				&github.SearchOptions{
					ListOptions: github.ListOptions{
						PerPage: 100,
						Page:    page,
					}})
			if err != nil {
				panic(err)
			}

			page = resp.NextPage

			if *result.IncompleteResults {
				return nil, errIncompleteResult
			}

			for _, v := range result.Commits {
				m[*v.Repository.FullName] += 1
			}

			if resp.NextPage == 0 {
				break
			}
		}
	}
	return m, nil
}
