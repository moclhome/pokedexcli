package main

import (
	"bootdev/go/pokedexcli/internal"
	"maps"
	"slices"
)

func commandHelp(c *internal.Config, param string) error {
	c.Printer.Printf("Usage:\n\n")
	for _, key := range slices.Sorted(maps.Keys(internal.Registry)) {
		command := internal.Registry[key]
		c.Printer.Printf("%s: %s\n", command.Name, command.Description)
	}
	c.Printer.Println("\nThe previous commands can be repeated by using the up arrow.")
	c.Printer.Println("Tab can be used for autocompletion on commands, areas that have lately been listed\nand Pokémon that have lately been encountered while exploring or that are available in your Pokedex.")
	return nil
}
