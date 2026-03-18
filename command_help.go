package main

import (
	"maps"
	"slices"
)

func commandHelp(c *config, param string) error {
	c.printer.Printf("Usage:\n\n")
	for _, key := range slices.Sorted(maps.Keys(registry)) {
		command := registry[key]
		c.printer.Printf("%s: %s\n", command.name, command.description)
	}
	c.printer.Printf("\nThe previous commands can be repeated by using the up arrow.\n")
	return nil
}
