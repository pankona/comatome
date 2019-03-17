package main

import "time"

// optStrtoTime converts yyyy-mm string to time.Time.
// parameter is a string that has layout like "2006-01".
func optStrtoTime(s string) (time.Time, error) {
	return time.Parse("2006-01", s)
}
