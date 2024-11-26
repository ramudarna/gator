package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	//Delete all users
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}

	fmt.Println("Database reset successful")
	return nil
}