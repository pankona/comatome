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
	f := resolveFlags()
	ft, err := fromto(f.from, f.to)
	if err != nil {
		panic(err)
	}

	c := comatome.NewClient(os.Getenv("GITHUB_API_TOKEN"))

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

func resolveFlags() flags {
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

	return f
}

// fromto function's behavior is depends on specified from and to.
//
// Legends:
//   - : not specified
//   o : specified
//
// | from | to | treat as
// |------|----|---------------------------------------------
// | -    | -  | from: 1 month ago from to, to: now
// | o    | -  | from: as specified,        to: now
// | -    | o  | from: 1 month ago from to, to: as specified
// | o    | o  | from: as specified,        to: as specified
//
// If from points future than to, this function returns error.
func fromto(from, to string) (*comatome.FromTo, error) {
	switch {
	case from == "" && to == "":
		t := time.Now()
		return comatome.NewFromTo(t.AddDate(0, -1, 0), t, time.Local), nil

	case from != "" && to == "":
		f, err := optStrtoTime(from)
		if err != nil {
			return nil, err
		}
		t := time.Now()
		return comatome.NewFromTo(f, t, time.Local), nil

	case from == "" && to != "":
		t, err := optStrtoTime(to)
		if err != nil {
			return nil, err
		}
		t = t.AddDate(0, 1, 0)
		t = t.AddDate(0, 0, -1)
		f := t.AddDate(0, -1, 0)
		return comatome.NewFromTo(f, t, time.Local), nil

	case from != "" && to != "":
		fallthrough

	default:
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
		if f.After(t) {
			return nil, fmt.Errorf("from must be past of to")
		}
		return comatome.NewFromTo(f, t, time.Local), nil
	}
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
