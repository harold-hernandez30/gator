package commands

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)


func CommandRegister(s *State, cmd Command) error {
	if len(cmd.Args) <= 0 {
		return fmt.Errorf("name is required")
	} 

	context := context.Background()
	
	name := cmd.Args[0]

	_, err := s.DB.GetUser(context, name)

	if err == nil {
		return fmt.Errorf("user with '%v' already exists", name)
	}

	params := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
	}

	user, err := s.DB.CreateUser(context, params)

	if err != nil {
		return err
	}

	fmt.Printf("New user created: %+v\n", user)
	s.Config.SetUser(user.Name)

	return nil
 }