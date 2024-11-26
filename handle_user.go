package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ramudarna/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}
	name := cmd.Args[0]

	//Check if the current user already exists
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("user with %s name does not exist", name)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("User switched successfully to : %s\n!", name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	//Check if the current user already exists
	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		return fmt.Errorf("user with %s name already exists", name)
	}
	//Create a new user
	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	createdUser, err := s.db.CreateUser(context.Background(), newUser)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(createdUser.Name)
	if err != nil {
		return err
	}

	fmt.Printf("User created: %v\n", createdUser)
	return nil
}

func handlerUsers(s *state, cmd command) error {
	//Get all users
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
