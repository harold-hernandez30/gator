package commands

import "fmt"


func CommandLogin(s *State, cmd Command) error {
	if len(cmd.Args) <= 0 {
		return fmt.Errorf("username is required")
	} 

	
	s.Config.SetUser(cmd.Args[0])
	fmt.Printf("New user logged in: %s\n", s.Config.CurrentUserName)

	return nil
 }