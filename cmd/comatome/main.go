package main

import (
	"flag"
	"fmt"
	"os"

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
	var fromto string

	flag.StringVar(&fromto, "fromto", "", "specify year and month by yyyymm format to fetch contributions")
	f := flags{}
	flag.BoolVar(&f.all, "all", false, "show all")
	flag.BoolVar(&f.commits, "co", false, "show commits")
	flag.BoolVar(&f.createdRepos, "re", false, "show created repositories")
	flag.BoolVar(&f.openedPRs, "op", false, "show opened pull requests")
	flag.BoolVar(&f.reviewedPRs, "rp", false, "show reviewed pull requests")
	flag.BoolVar(&f.openedIssues, "oi", false, "show opened issues")

	flag.Parse()

	if fromto == "" {
		fmt.Println("fromto must be specified")
		os.Exit(1)
	}
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

	if f.commits {
		commits(c, fromto)
		fmt.Println()
	}
	if f.createdRepos {
		createdRepos(c, fromto)
		fmt.Println()
	}
	if f.openedPRs {
		openedPullRequests(c, fromto)
		fmt.Println()
	}
	if f.reviewedPRs {
		reviewedPullRequests(c, fromto)
		fmt.Println()
	}
	if f.openedIssues {
		openedIssues(c, fromto)
		fmt.Println()
	}
}

func commits(c *comatome.Client, fromto string) {
	results, err := comatome.QueryCommitsPerRepo(c, fromto)
	if err != nil {
		panic(err)
	}
	comatome.ShowCommitsPerRepo(results)
}

func createdRepos(c *comatome.Client, fromto string) {
	results, err := comatome.QueryCreatedRepos(c, fromto)
	if err != nil {
		panic(err)
	}
	comatome.ShowCreatedRepos(results)
}

func openedPullRequests(c *comatome.Client, fromto string) {
	results, err := comatome.QueryOpenedPullRequests(c, fromto)
	if err != nil {
		panic(err)
	}
	comatome.ShowOpenedPullRequests(results)
}

func reviewedPullRequests(c *comatome.Client, fromto string) {
	results, err := comatome.QueryReviewedPullRequests(c, fromto)
	if err != nil {
		panic(err)
	}
	comatome.ShowReviewedPullRequests(results)
}

func openedIssues(c *comatome.Client, fromto string) {
	results, err := comatome.QueryOpenedIssues(c, fromto)
	if err != nil {
		panic(err)
	}
	comatome.ShowOpenedIssues(results)
}
