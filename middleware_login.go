package main

import (
	"context"
	"fmt"
	"github.com/mcdotjs/blog_aggregator/internal/database"
	"os"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		context := context.Background()
		user, err := s.db.GetUser(context, s.cfg.CurrentUserName)
		if err != nil {
			fmt.Println(s.cfg.CurrentUserName + " neni v db")
			os.Exit(1)
		}
		return handler(s, cmd, user)
	}
}
