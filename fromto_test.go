package comatome

import (
	"testing"
	"time"
)

func loadLocation(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}
	return loc
}

func timeParse(layout, target string) time.Time {
	t, err := time.Parse(layout, target)
	if err != nil {
		panic(err)
	}
	return t
}

func TestQueryStr(t *testing.T) {
	tcs := []struct {
		inFrom   time.Time
		inTo     time.Time
		inLoc    *time.Location
		wantFrom string
		wantTo   string
	}{
		// only yyyy-mm is acceptable
		{
			inFrom:   timeParse("2006-01", "2019-01"),
			inTo:     timeParse("2006-01", "2019-02"),
			inLoc:    loadLocation("Asia/Tokyo"),
			wantFrom: "2019-01-01T09:00:00+09:00",
			wantTo:   "2019-02-01T09:00:00+09:00",
		},
	}

	for i, tc := range tcs {
		ft := NewFromTo(tc.inFrom, tc.inTo, tc.inLoc)
		gotFrom, gotTo := ft.QueryStr()

		if gotFrom != tc.wantFrom {
			t.Fatalf("[No.%d] unexpected result. [got] %v [want] %v", i, gotFrom, tc.wantFrom)
		}
		if gotTo != tc.wantTo {
			t.Fatalf("[No.%d] unexpected result. [got] %v [want] %v", i, gotTo, tc.wantTo)
		}
	}
}
