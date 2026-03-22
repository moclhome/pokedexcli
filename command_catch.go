package main

import (
	"math/rand"
	"strings"
)

func commandCatch(c *config, param string) error {
	if _, found := c.caughtPokemons[param]; found {
		c.printer.Printf("You already have %s in your Pokédex. Try to catch someone else!\n", param)
		return nil
	}
	thePokemon, err := c.pokeapiClient.GetPokemon(param, c.flagDebug)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			c.printer.Printf("There is no Pokémon named %s.\n", param)
			return nil
		}
		return err
	}
	c.printer.Printf("Throwing a Pokeball at %s...\n", param)
	dice := rand.Intn(250)
	c.printer.Printf("%s has experience %d. Throwing the dice... It was %d.\n", param, thePokemon.BaseExperience, dice)
	if dice > thePokemon.BaseExperience {
		c.printer.Printf("You caught %s!\n", param)
		c.caughtPokemons[param] = thePokemon
		c.printer.Printf("There are now %d Pokémon in your Pokedex.\n", len(c.caughtPokemons))
		// save caught pokemon for autocompletion
		CurrentCompletionData["pokemon"] = append(CurrentCompletionData["pokemon"], param)
	} else {
		c.printer.Printf("%s escaped.\n", param)
	}
	return nil
}
