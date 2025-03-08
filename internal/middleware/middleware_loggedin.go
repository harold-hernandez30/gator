package middleware

import (
	"context"
	"gator/internal/commands"
	"gator/internal/database"
)

func MiddlewareHandleLoggedIn(
	handler func(state *commands.State, cmd commands.Command, user database.User) error) func(state *commands.State, cmd commands.Command) error {

	return func (s *commands.State, cmd commands.Command) error {
	
		currentUserName := s.Config.CurrentUserName
		user, err := s.DB.GetUser(context.Background(), currentUserName)

		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}

func IsLoggedIn(state *commands.State, cmd commands.Command) error {
	
	currentUserName := state.Config.CurrentUserName
	_, err := state.DB.GetUser(context.Background(), currentUserName)

	if err != nil {
		return err
	}

	return nil
}