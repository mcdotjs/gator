package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/mcdotjs/blog_aggregator/internal/database"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func addFeed(s *state, cmd command, user database.User) error {
	fmt.Println("addFeed", cmd.Name, user)
	if len(cmd.Args) < 2 {
		err := fmt.Errorf("We want two arguments")
		return err
	}
	context := context.Background()

	newfeed := &database.CreateFeedParams{
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context, *newfeed)
	if err != nil {
		return err
	}

	newFoolowFeed := &database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		UserID:    user.ID,
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
		fmt.Println(cmd.Name, "error")
		return err
	}
	for _, feed := range feeds {

		user, err := s.db.GetUserById(context, feed.UserID)
		if err != nil {
			return err
		}
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(user.Name)
	}
	return nil
}

func scrapeFeeds(s *state, cmd command) error {
	context := context.Background()
	user, err := s.db.GetUser(context, s.cfg.CurrentUserName)
	fmt.Println(user.Name)
	if err != nil {
		return fmt.Errorf("User neni v db: %w", err)
	}

	lastFeed, err := s.db.GetNextFeedToFetch(context, user.ID)
	if err != nil {
		return fmt.Errorf("problem getNextFeed: %w", err)
	}

	params := &database.MarkFeedFetchedParams{
		ID:     lastFeed.ID,
		UserID: user.ID,
	}

	feed, err := s.db.MarkFeedFetched(context, *params)
	if err != nil {
		return fmt.Errorf("Problem whith MarkFeedFetched: %w", err)
	}
	log.Printf("Scraping feed: %s for user: %s", feed.Url, user.Name)

	feeds, err := fetchFeed(context, feed.Url)
	if err != nil {
		return fmt.Errorf("Problem fetch feed by url: %w", err)
	}

	for _, v := range feeds.Channel.Item {
		//NOTE:
		// Convert string to sql.NullString
		description := sql.NullString{
			String: v.Description,
			Valid:  v.Description != "", // Valid if non-empty
		}

		//NOTE:
		// Parse v.PubDate into time.Time, and convert to sql.NullTime
		var publishedAt sql.NullTime
		if pubTime, err := time.Parse(time.RFC1123, v.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  pubTime,
				Valid: true,
			}
		} else {
			//NOTE:
			// If parsing fails, leave publishedAt as invalid
			publishedAt = sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			}
		}
		postCreateParams := &database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       v.Title,
			Url:         v.Link,
			FeedID:      feed.ID,
			Description: description,
			PublishedAt: publishedAt,
		}

		_, err := s.db.CreatePost(context, *postCreateParams)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}

	}
	return nil
}

func browsePosts(s *state, cmd command, user database.User) error {
	context := context.Background()
	limit := 2
	if len(cmd.Args) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = specifiedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}
	log.Printf("User %d requested to browse posts with limit %d", user.Name, limit)
	params := &database.GetPostsForUserTroughJoinParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.db.GetPostsForUserTroughJoin(context, *params)
	if err != nil {
		return fmt.Errorf("browse posts error: %w", err)
	}

	for _, p := range posts {
		fmt.Println(p.Title)
	}
	return nil
}
