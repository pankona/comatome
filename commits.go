package main

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/google/go-github/v21/github"
)

type CommitsPerRepo map[string]int

var errIncompleteResult = errors.New("incomplete result error")

const fromto = "2018-10-01..2018-12-31"

func QueryCommitsPerRepo(c *Client) (CommitsPerRepo, error) {
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
		m, err = queryCommitsPerRepo(c, emails)
		if err == errIncompleteResult {
			<-time.After(1 * time.Second)
			continue
		}
		break
	}

	return m, err
}

func queryCommitsPerRepo(c *Client, emails []*github.UserEmail) (CommitsPerRepo, error) {

	m := make(CommitsPerRepo)

	for _, email := range emails {
		page := 1
		for {
			result, resp, err := c.Search.Commits(
				context.Background(),
				fmt.Sprintf("author-email:%s+author-date:%s", email.GetEmail(), fromto),
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

func ShowCommitsPerRepo(m CommitsPerRepo) {
	keys := make([]string, len(m))
	index := 0
	for k, _ := range m {
		keys[index] = k
		index++
	}
	sort.Strings(keys)

	total := 0
	fmt.Printf("commits on %s\n", fromto)
	for _, v := range keys {
		fmt.Printf("%3d commits on %s\n", m[v], v)
		total += m[v]
	}
	fmt.Printf("total %d commits\n", total)
}
