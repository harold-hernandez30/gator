package commands

import (
	"context"
	"fmt"
)

func CommandGetUsers (state *State, cmd Command) error {
	users, err := state.DB.GetUsers(context.Background())

	if err != nil {
		return err
	}


	for _, user := range users {

		if user.Name.String == state.Config.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name.String)
		} else {
			fmt.Printf("* %s\n", user.Name.String)
		}
	}
	return nil
}