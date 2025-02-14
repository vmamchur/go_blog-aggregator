package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vmamchur/go_blog-aggregator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("Couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("Couldn't set current user: %w", err)
	}

	fmt.Printf("User %s switched successfully\n", name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name})
	if err != nil {
		return fmt.Errorf("Couldn't create user: %w", err)
	}
	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("Couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func handlerListUsers(s *state, cmd command) error {
    users, err := s.db.GetUsers(context.Background())
    if err != nil {
        return fmt.Errorf("Couldn't list users: %w", err)
    }

    for _, user := range users {
        if user.Name == s.cfg.CurrentUserName {
            fmt.Printf("* %v (current)\n", user.Name)
            continue
        }
        fmt.Printf("* %v\n", user.Name)
    }

    return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:    %v\n", user.ID)
	fmt.Printf(" * Name:  %v\n", user.Name)
}
