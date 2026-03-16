package main

import (
	"fmt"
)

func commandInspect(c *config, param string) error {
	if pokemon, found := c.caughtPokemons[param]; !found {
		fmt.Printf("You have not yet caught the Pokémon %s.\n", param)
	} else {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for i := 0; i < len(pokemon.Types); i++ { // why is this not working with range?
			fmt.Printf("  - %s\n", pokemon.Types[i].Type.Name)
		}
		/*for _, type := range pokemon.Types { // I get the error "expected 1 expression"
			fmt.Printf("  - %s\n", type.Type.Name)
		}*/
	}
	return nil
}
