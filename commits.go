package main

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/google/go-github/v21/github"
)

// CommitsPerRepo is a map to represent commits/repository
type CommitsPerRepo map[string]int

var errIncompleteResult = errors.New("incomplete result error")

// QueryCommitsPerRepo queries commits per repository created on specified term (fromto)
func QueryCommitsPerRepo(c *Client, fromto string) (CommitsPerRepo, error) {
	var (
		m        CommitsPerRepo
		err      error
		maxRetry = 5
	)

	emails, _, err := c.Users.ListEmails(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	for i := 0; i < maxRetry; i++ {
		m, err = queryCommitsPerRepo(c, emails, fromto)
		if err == errIncompleteResult {
			<-time.After(1 * time.Second)
			continue
		}
		break
	}

	return m, err
}

func queryCommitsPerRepo(c *Client, emails []*github.UserEmail, fromto string) (CommitsPerRepo, error) {
	m := make(CommitsPerRepo)

	for _, email := range emails {
		page := 1
		for {
			// Notice:
			// if fromto is 2018-01-01..2018-01-31, the result doesn't include 2018-01-01's commits.
			// if 2018-01-01's commits should be included, fromto should be like 2018-12-31..2019-01-31
			result, resp, err := c.Search.Commits(
				context.Background(),
				fmt.Sprintf("author-email:%s author-date:%s", email.GetEmail(), fromto),
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
				m[*v.Repository.FullName]++
			}

			if resp.NextPage == 0 {
				break
			}
		}
	}

	return m, nil
}

// ShowCommitsPerRepo shows commits/repositories
func ShowCommitsPerRepo(m CommitsPerRepo) {
	keys := make([]string, len(m))
	index := 0
	for k := range m {
		keys[index] = k
		index++
	}
	sort.Strings(keys)

	total := 0
	for _, v := range keys {
		total += m[v]
	}

	fmt.Printf("Created %d commits in %d repositories\n", total, len(m))
	for _, v := range keys {
		fmt.Printf("%d\t%s\n", m[v], v)
	}
}
