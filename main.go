package main

import (
	"flag"
	"os"
)

func main() {
	fromto := flag.String("fromto", "", "specify year and month by yyyymm format to fetch contributions")
	flag.Parse()

	if *fromto == "" {
		panic("fromto is not specified.")
	}
	c := NewClient(os.Getenv("GITHUB_API_TOKEN"))

	m, err := QueryCommitsPerRepo(c, *fromto)
	if err != nil {
		panic(err)
	}

	ShowCommitsPerRepo(m)
}
