package main

import (
	"context"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("error handler login")
	}

	passedName := cmd.Args[0]
	context := context.Background()
	_, err := s.db.GetUser(context, passedName)
	if err != nil {
		fmt.Println(passedName + " neni v db")
		os.Exit(1)
	}
	s.cfg.SetUser(cmd.Args[0])
	return nil
}
