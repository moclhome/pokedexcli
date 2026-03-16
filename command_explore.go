package main

import (
	"fmt"
)

func commandExplore(c *config, param string) error {
	fmt.Printf("Exploring %s ...\n", param)
	if err := c.pokeapiClient.PrintPokemonEncounters(param, c.flagDebug); err != nil {
		return err
	}
	return nil
}
