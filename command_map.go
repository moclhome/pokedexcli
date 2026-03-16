package main

import (
	"fmt"
)

func commandMap(c *config, param string) error {
	if c.Next == "" {
		fmt.Println("you're on the last page")
		return nil
	}

	locAreas, err := c.pokeapiClient.PrintLocationAreas(c.Next, c.flagDebug)
	if err != nil {
		return err
	}

	c.Next = locAreas.Next
	c.Previous = locAreas.Previous
	return nil
}

func commandMapb(c *config, param string) error {
	if c.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locAreas, err := c.pokeapiClient.PrintLocationAreas(c.Previous, c.flagDebug)
	if err != nil {
		return err
	}

	c.Next = locAreas.Next
	c.Previous = locAreas.Previous
	return nil
}
