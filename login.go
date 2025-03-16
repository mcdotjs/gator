package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("error handler login")
	}

	s.value.SetUser(cmd.Args[0])
	return nil
}
