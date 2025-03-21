package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mcdotjs/blog_aggregator/internal/database"
	"os"
	"time"
)

func handlerRegister(s *state, cmd command) error {
	passedName := cmd.Args[0]
	context := context.Background()
	p := &database.CreateUserParams{
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		Name:      passedName,
	}
	_, err := s.db.GetUser(context, passedName)
	if err == nil {
		fmt.Println(passedName + " already exist")
		os.Exit(1)
	}

	s.cfg.SetUser(passedName)
	s.db.CreateUser(context, *p)
	fmt.Println("register llll", passedName)
	return nil
}
