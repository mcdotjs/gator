package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mcdotjs/blog_aggregator/internal/database"
	"os"
	"time"
)

func addFeed(s *state, cmd command) error {

	fmt.Println("args", cmd.Args)
	if len(cmd.Args) < 2 {
		err := fmt.Errorf("We want two arguments")
		os.Exit(1)
		return err
	}
	context := context.Background()

	logedUser, err := s.db.GetUser(context, s.cfg.CurrentUserName)
	if err != nil {
		os.Exit(1)
		return err
	}
	fmt.Println(logedUser)
	newfeed := &database.CreateFeedParams{
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    logedUser.ID,
	}
	s.db.CreateFeed(context, *newfeed)
	os.Exit(0)
	return nil
}
