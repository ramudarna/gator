package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ramudarna/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	url := cmd.Args[0]
	feed, err := s.db.GetFeedFromUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error fetching feeds: %w", err)
	}
	//Create a new feed follow for the current user
	ffRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feef_follow: %w", err)
	}
	fmt.Println("Feed follow created:")
	printFeedFollow(ffRow.UserName, ffRow.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error fetching feed follows: %w", err)
	}
	if len(feedFollows) == 0 {
		fmt.Println("No feed follows for this user")
		return nil
	}
	fmt.Printf("Feed follows for user: %s\n", user.Name)
	for _, ff := range feedFollows {
		fmt.Printf("feed name: %s\n", ff.FeedName)
	}
	return nil
}

func handlerUnFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	url := cmd.Args[0]
	feed, err := s.db.GetFeedFromUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}
	err = s.db.DeleteFeedFollowsByUserFeedUrl(context.Background(), database.DeleteFeedFollowsByUserFeedUrlParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error deleting feed follows: %w", err)
	}
	fmt.Println("Unfollowed successfully")
	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
