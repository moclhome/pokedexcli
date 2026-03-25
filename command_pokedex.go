package main

import "bootdev/go/pokedexcli/internal"

func commandPokedex(c *internal.Config, param string) error {
	if len(c.CaughtPokemons) == 0 {
		c.Printer.Println("Your Pokedex is still empty. Try to catch a Pokémon!")
	} else {
		c.Printer.Printf("There are %d Pokémon in your Pokedex:\n", len(c.CaughtPokemons))
		for _, thePokemon := range c.CaughtPokemons {
			c.Printer.Printf("  -%s\n", thePokemon.Name)
		}
	}
	return nil
}
