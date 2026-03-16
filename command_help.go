package main

import (
	"fmt"
	"maps"
	"slices"
)

func commandHelp(c *config, param string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, key := range slices.Sorted(maps.Keys(registry)) {
		command := registry[key]
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
