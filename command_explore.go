package main

import "strings"

func commandExplore(c *config, param string) error {
	c.printer.Printf("Exploring %s ...\n", param)
	theAreaWithEncounter, err := c.pokeapiClient.GetPokemonEncounter(param, c.flagDebug, c.printer)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			c.printer.Printf("There is no area named %s.\n", param)
			return nil
		}
		return err
	}
	for _, encounter := range theAreaWithEncounter.PokemonEncounters {
		c.printer.Printf(" - %s\n", encounter.Pokemon.Name)
		// save pokemon encounters for autocompletion
		CurrentCompletionData["pokemon"] = append(CurrentCompletionData["pokemon"], encounter.Pokemon.Name)
	}
	return nil
}
