package main

func commandExplore(c *config, param string) error {
	c.printer.Printf("Exploring %s ...\n", param)
	if err := c.pokeapiClient.PrintPokemonEncounters(param, c.flagDebug, c.printer); err != nil {
		return err
	}
	return nil
}
