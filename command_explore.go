package main

import (
	"bootdev/go/pokedexcli/internal"
	"strings"
)

func commandExplore(c *internal.Config, param string) error {
	c.Printer.Printf("Exploring %s ...\n", param)
	theAreaWithEncounter, err := c.PokeapiClient.GetPokemonEncounter(param, c.FlagDebug)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			c.Printer.Printf("There is no area named %s.\n", param)
			return nil
		}
		return err
	}
	c.Printer.Println("Pokemon found:")
	for _, encounter := range theAreaWithEncounter.PokemonEncounters {
		c.Printer.Printf(" - %s\n", encounter.Pokemon.Name)
		// save pokemon encounters for autocompletion
		CurrentCompletionData["pokemon"] = append(CurrentCompletionData["pokemon"], encounter.Pokemon.Name)
	}
	return nil
}
