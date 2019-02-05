package main

import (
	"context"
)

// Username returns username by specified GitHub API Token
func Username(c *Client) string {
	user, _, err := c.Users.Get(context.Background(), "")
	if err != nil {
		panic(err)
	}
	return user.GetLogin()
}
