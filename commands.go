package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	value map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) error {
	c.value[name] = f
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	fmt.Println(cmd)
	handler, exists := c.value[cmd.Name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}
