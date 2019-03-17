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

	from, to string
}

func main() {
	f := flags{}
	flag.BoolVar(&f.all, "all", false, "show all")
	flag.BoolVar(&f.commits, "co", false, "show commits")
	flag.BoolVar(&f.createdRepos, "re", false, "show created repositories")
	flag.BoolVar(&f.openedPRs, "op", false, "show opened pull requests")
	flag.BoolVar(&f.reviewedPRs, "rp", false, "show reviewed pull requests")
	flag.BoolVar(&f.openedIssues, "oi", false, "show opened issues")
	flag.StringVar(&f.from, "from", "", "specify time contributed from (yyyy-mm)")
	flag.StringVar(&f.to, "to", "", "specify time contributed to (yyyy-mm)")
	flag.Parse()

	// if flag is not specified, treat as all
	if !f.commits &&
		!f.createdRepos &&
		!f.openedPRs &&
		!f.reviewedPRs &&
		!f.openedIssues {
		f.all = true
	}

	if f.all {
		f = flags{
			commits:      true,
			createdRepos: true,
			openedPRs:    true,
			reviewedPRs:  true,
			openedIssues: true,
		}
	}

	c := comatome.NewClient(os.Getenv("GITHUB_API_TOKEN"))
	ft, err := fromto(f.from, f.to)
	if err != nil {
		panic(err)
	}

	if f.commits {
		commits(c, ft)
		fmt.Println()
	}
	if f.createdRepos {
		createdRepos(c, ft)
		fmt.Println()
	}
	if f.openedPRs {
		openedPullRequests(c, ft)
		fmt.Println()
	}
	if f.reviewedPRs {
		reviewedPullRequests(c, ft)
		fmt.Println()
	}
	if f.openedIssues {
		openedIssues(c, ft)
		fmt.Println()
	}
}

func fromto(from, to string) (*comatome.FromTo, error) {
	if from != "" || to != "" {
		f, err := optStrtoTime(from)
		if err != nil {
			return nil, err
		}
		t, err := optStrtoTime(to)
		if err != nil {
			return nil, err
		}
		t = t.AddDate(0, 1, 0)
		t = t.AddDate(0, 0, -1)
		return comatome.NewFromTo(f, t, nil), nil
	}

	// one month ago from now
	now := time.Now()
	return comatome.NewFromTo(now.AddDate(0, -1, 0), now, nil), nil
}

func commits(c *comatome.Client, fromto *comatome.FromTo) {
	results, err := comatome.QueryCommitsPerRepo(c, fromto)
	if err != nil {
		panic(err)
	}
	comatome.ShowCommitsPerRepo(results)
}

func createdRepos(c *comatome.Client, fromto *comatome.FromTo) {
	results, err := comatome.QueryCreatedRepos(c, fromto)
	if err != nil {
		panic(err)
	}
	comatome.ShowCreatedRepos(results)
}

func openedPullRequests(c *comatome.Client, fromto *comatome.FromTo) {
	results, err := comatome.QueryOpenedPullRequests(c, fromto)
	if err != nil {
		panic(err)
	}
	comatome.ShowOpenedPullRequests(results)
}

func reviewedPullRequests(c *comatome.Client, fromto *comatome.FromTo) {
	results, err := comatome.QueryReviewedPullRequests(c, fromto)
	if err != nil {
		panic(err)
	}
	comatome.ShowReviewedPullRequests(results)
}

func openedIssues(c *comatome.Client, fromto *comatome.FromTo) {
	results, err := comatome.QueryOpenedIssues(c, fromto)
	if err != nil {
		panic(err)
	}
	comatome.ShowOpenedIssues(results)
}
