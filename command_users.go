package main

import (
	"context"
	"fmt"
	"os"
)

func handlerGetUsers(s *state, cmd command) error {

	context := context.Background()
	users, err := s.db.GetAllUsers(context)
	if err != nil {
		os.Exit(1)
		fmt.Println(cmd.Name, "error")
		return err
	}
	for _, u := range users {
		if s.cfg.CurrentUserName == u.Name {
			fmt.Printf("* %s (current)\n", u.Name)
		} else {
			fmt.Printf("* %s \n", u.Name)
		}
	}

	os.Exit(0)
	return nil
}
