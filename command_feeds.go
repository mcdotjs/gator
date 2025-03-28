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

	newfeed := &database.CreateFeedParams{
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    logedUser.ID,
	}

	feed, err := s.db.CreateFeed(context, *newfeed)
	if err != nil {
		os.Exit(1)
		return err
	}

	newFoolowFeed := &database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		UserID:    logedUser.ID,
		FeedID:    feed.ID,
	}

	s.db.CreateFeedFollow(context, *newFoolowFeed)
	os.Exit(0)
	return nil
}

func getFeeds(s *state, cmd command) error {
	context := context.Background()
	feeds, err := s.db.GetFeeds(context)
	if err != nil {
		os.Exit(1)
		fmt.Println(cmd.Name, "error")
		return err
	}
	for _, feed := range feeds {

		user, err := s.db.GetUserById(context, feed.UserID)
		if err != nil {
			os.Exit(1)
			return err
		}
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(user.Name)
	}
	return nil
}
