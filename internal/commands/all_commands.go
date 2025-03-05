package commands

import "fmt"

type Commands struct {
	Registered map[string]func(state *State, cmd Command) error
}

func (c *Commands) Register(name string, f func(state *State, cmd Command) error) {
	c.Registered[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	registeredCommand, found := c.Registered[cmd.Name]

	if found {
		return registeredCommand(s, cmd)
	} else {
		return fmt.Errorf("command '%s' not supported", cmd.Name)
	}
}