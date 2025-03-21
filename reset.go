package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {
	context := context.Background()
	err := s.db.DeleteAllUsers(context)
	if err != nil {
		fmt.Println(cmd.Name, " was successful.")
		os.Exit(1)
		return err
	}

	os.Exit(0)
	return nil
}
