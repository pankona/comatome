package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/pankona/comatome"
)

type flags struct {
	all          bool
	commits      bool
	createdRepos bool
	openedPRs    bool
	reviewedPRs  bool
	openedIssues bool
}

func main() {
	f := flags{}
	flag.BoolVar(&f.all, "all", false, "show all")
	flag.BoolVar(&f.commits, "co", false, "show commits")
	flag.BoolVar(&f.createdRepos, "re", false, "show created repositories")
	flag.BoolVar(&f.openedPRs, "op", false, "show opened pull requests")
	flag.BoolVar(&f.reviewedPRs, "rp", false, "show reviewed pull requests")
	flag.BoolVar(&f.openedIssues, "oi", false, "show opened issues")

	flag.Parse()

	c := comatome.NewClient(os.Getenv("GITHUB_API_TOKEN"))

	if f.all {
		f = flags{
			commits:      true,
			createdRepos: true,
			openedPRs:    true,
			reviewedPRs:  true,
			openedIssues: true,
		}
	}

	var (
		now = time.Now()

		from = now.AddDate(0, -1, 0)
		to   = now
	)

	if f.commits {
		commits(c, from, to)
		fmt.Println()
	}
	if f.createdRepos {
		createdRepos(c, from, to)
		fmt.Println()
	}
	if f.openedPRs {
		openedPullRequests(c, from, to)
		fmt.Println()
	}
	if f.reviewedPRs {
		reviewedPullRequests(c, from, to)
		fmt.Println()
	}
	if f.openedIssues {
		openedIssues(c, from, to)
		fmt.Println()
	}
}

func commits(c *comatome.Client, from, to time.Time) {
	results, err := comatome.QueryCommitsPerRepo(c, from, to)
	if err != nil {
		panic(err)
	}
	comatome.ShowCommitsPerRepo(results)
}

func createdRepos(c *comatome.Client, from, to time.Time) {
	results, err := comatome.QueryCreatedRepos(c, from, to)
	if err != nil {
		panic(err)
	}
	comatome.ShowCreatedRepos(results)
}

func openedPullRequests(c *comatome.Client, from, to time.Time) {
	results, err := comatome.QueryOpenedPullRequests(c, from, to)
	if err != nil {
		panic(err)
	}
	comatome.ShowOpenedPullRequests(results)
}

func reviewedPullRequests(c *comatome.Client, from, to time.Time) {
	results, err := comatome.QueryReviewedPullRequests(c, from, to)
	if err != nil {
		panic(err)
	}
	comatome.ShowReviewedPullRequests(results)
}

func openedIssues(c *comatome.Client, from, to time.Time) {
	results, err := comatome.QueryOpenedIssues(c, from, to)
	if err != nil {
		panic(err)
	}
	comatome.ShowOpenedIssues(results)
}
