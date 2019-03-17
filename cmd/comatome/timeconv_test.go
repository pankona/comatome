package main

import (
	"testing"
	"time"
)

func timeParse(layout, target string) time.Time {
	t, err := time.Parse(layout, target)
	if err != nil {
		panic(err)
	}
	return t
}

func TestOptStrtoTime(t *testing.T) {
	tcs := []struct {
		inStr      string
		wantTime   time.Time
		wantErrNil bool
	}{
		// only yyyy-mm is acceptable
		{
			inStr:      "2019-01",
			wantTime:   timeParse("2006-01", "2019-01"),
			wantErrNil: true,
		},

		// other layout is not acceptable
		{
			inStr:      "2019-01-01",
			wantTime:   time.Time{},
			wantErrNil: false,
		},
		{
			inStr:      "2019-01-01 15:04:05",
			wantTime:   time.Time{},
			wantErrNil: false,
		},
	}

	for i, tc := range tcs {
		ti, err := optStrtoTime(tc.inStr)
		if tc.wantErrNil && err != nil {
			t.Fatalf("[No.%d] unexpected err. [got] %v [want] nil", i, err)
		}
		if ti != tc.wantTime {
			t.Fatalf("[No.%d] unexpected result. [got] %v [want] %v", i, ti, tc.wantTime)
		}
	}
}
