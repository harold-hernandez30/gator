package commands

import (
	"context"
	"database/sql"
	"fmt"
)


func CommandLogin(s *State, cmd Command) error {
	if len(cmd.Args) <= 0 {
		return fmt.Errorf("username is required")
	} 

	user, err := s.DB.GetUser(context.Background(), sql.NullString{
		String: cmd.Args[0],
		Valid: true,
	})

	if err != nil {
		return fmt.Errorf("username: %s, not found", cmd.Args[0])
	}
	
	s.Config.SetUser(user.Name.String)
	fmt.Printf("New user logged in: %s\n", user.Name.String)

	return nil
 }