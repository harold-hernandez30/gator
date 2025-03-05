package commands

import (
	"context"
	"database/sql"
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
	
	nameNullString := sql.NullString{
		String: cmd.Args[0],
		Valid: true,
	}

	_, err := s.DB.GetUser(context, nameNullString)

	if err == nil {
		return fmt.Errorf("user with '%v' already exists", nameNullString.String)
	}

	params := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: nameNullString,
	}

	user, err := s.DB.CreateUser(context, params)

	if err != nil {
		return err
	}

	fmt.Printf("New user created: %+v\n", user)
	s.Config.SetUser(user.Name.String)

	return nil
 }