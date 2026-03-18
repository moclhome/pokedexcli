package main

func commandPokedex(c *config, param string) error {
	if len(c.caughtPokemons) == 0 {
		c.printer.Println("Your Pokedex is still empty. Try to catch a Pokémon!")
	} else {
		c.printer.Printf("There are %d Pokémon in your Pokedex:\n", len(c.caughtPokemons))
		for _, thePokemon := range c.caughtPokemons {
			c.printer.Printf("  -%s\n", thePokemon.Name)
		}
	}
	return nil
}
