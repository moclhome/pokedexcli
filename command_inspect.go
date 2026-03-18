package main

func commandInspect(c *config, param string) error {
	if pokemon, found := c.caughtPokemons[param]; !found {
		c.printer.Printf("You have not yet caught the Pokémon %s.\n", param)
	} else {
		c.printer.Printf("Name: %s\n", pokemon.Name)
		c.printer.Printf("Height: %d\n", pokemon.Height)
		c.printer.Printf("Weight: %d\n", pokemon.Weight)
		c.printer.Println("Stats:")
		for _, stat := range pokemon.Stats {
			c.printer.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		c.printer.Println("Types:")
		for i := 0; i < len(pokemon.Types); i++ { // why is this not working with range?
			c.printer.Printf("  - %s\n", pokemon.Types[i].Type.Name)
		}
		/*for _, type := range pokemon.Types { // I get the error "expected 1 expression" TODO
			fmt.Printf("  - %s\n", type.Type.Name)
		}*/
	}
	return nil
}
