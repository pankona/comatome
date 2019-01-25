package main

import (
	"flag"
	"fmt"
	"os"
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
