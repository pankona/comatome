package comatome

import (
	"time"
)

// FromTo represents term of searching contributions
type FromTo struct {
	From time.Time
	To   time.Time
	Loc  *time.Location
}

// NewFromTo returns an instance of FromTo.
// Parameters (from, to) should be formed as yyy-mm.
// If specified parameter is not expected form,
// this function return error.
// loc is location. If nil is specified to loc, time.Local will be used.
func NewFromTo(from, to time.Time, loc *time.Location) *FromTo {
	return &FromTo{From: from, To: to, Loc: loc}
}

// QueryStr converts time.Time to string.
// returned string's layout is like "2006-01-02T15:04:05+09:00"
func (f *FromTo) QueryStr() (string, string) {
	l := f.Loc
	if l == nil {
		l = time.Local
	}
	return f.From.In(l).Format(time.RFC3339),
		f.To.In(l).Format(time.RFC3339)
}
