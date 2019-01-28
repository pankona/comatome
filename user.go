package main

import (
	"context"
)

func Username(c *Client) string {
	user, _, err := c.Users.Get(context.Background(), "")
	if err != nil {
		panic(err)
	}
	return user.GetLogin()
}
