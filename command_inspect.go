package main

import "bootdev/go/pokedexcli/internal"

func commandInspect(c *internal.Config, param string) error {
	if pokemon, found := c.CaughtPokemons[param]; !found {
		c.Printer.Printf("You have not yet caught the Pokémon %s.\n", param)
	} else {
		c.Printer.Printf("Name: %s\n", pokemon.Name)
		c.Printer.Printf("Height: %d\n", pokemon.Height)
		c.Printer.Printf("Weight: %d\n", pokemon.Weight)
		c.Printer.Println("Stats:")
		for _, stat := range pokemon.Stats {
			c.Printer.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		c.Printer.Println("Types:")
		for i := 0; i < len(pokemon.Types); i++ { // why is this not working with range?
			c.Printer.Printf("  - %s\n", pokemon.Types[i].Type.Name)
		}
		/*for _, type := range pokemon.Types { // I get the error "expected 1 expression" TODO
			fmt.Printf("  - %s\n", type.Type.Name)
		}*/
	}
	return nil
}
