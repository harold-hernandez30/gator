package commands

import (
	"context"
	"fmt"
)


func CommandLogin(s *State, cmd Command) error {
	if len(cmd.Args) <= 0 {
		return fmt.Errorf("username is required")
	} 

	user, err := s.DB.GetUser(context.Background(), cmd.Args[0])

	if err != nil {
		return fmt.Errorf("username: %s, not found", cmd.Args[0])
	}
	
	s.Config.SetUser(user.Name)
	fmt.Printf("New user logged in: %s\n", user.Name)

	return nil
 }