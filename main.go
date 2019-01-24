package main

import (
	"os"
)

func main() {
	c := NewClient(os.Getenv("GITHUB_API_TOKEN"))

	m, err := QueryCommitsPerRepo(c)
	if err != nil {
		panic(err)
	}

	ShowCommitsPerRepo(m)
}
