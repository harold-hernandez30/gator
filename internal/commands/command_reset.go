package commands

import (
	"context"
	"fmt"
)


func CommandReset(s *State, cmd Command) error {

	err := s.DB.DeleteAllRecords(context.Background())

	if err != nil {
		return fmt.Errorf("could not delete all user rows")
	}
	fmt.Print("removed all user rows\n")

	return nil
 }