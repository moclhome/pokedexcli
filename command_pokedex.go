package main

import (
	"fmt"
)

func commandPokedex(c *config, param string) error {
	if len(c.caughtPokemons) == 0 {
		fmt.Println("Your Pokedex is still empty. Try to catch a Pokémon!")
	} else {
		fmt.Printf("There are %d Pokémon in your Pokedex:\n", len(c.caughtPokemons))
		for _, thePokemon := range c.caughtPokemons {
			fmt.Printf("  -%s\n", thePokemon.Name)
		}
	}
	return nil
}
