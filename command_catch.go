package main

import (
	"bootdev/go/pokedexcli/internal"
	"math/rand"
	"strings"
)

func commandCatch(c *internal.Config, param string) error {
	if _, found := c.CaughtPokemons[param]; found {
		c.Printer.Printf("You already have %s in your Pokédex. Try to catch someone else!\n", param)
		return nil
	}
	thePokemon, err := c.PokeapiClient.GetPokemon(param, c.FlagDebug)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			c.Printer.Printf("There is no Pokémon named %s.\n", param)
			return nil
		}
		return err
	}
	c.Printer.Printf("Throwing a Pokeball at %s...\n", param)
	dice := rand.Intn(250)
	c.Printer.Printf("%s has experience %d. Throwing the dice... It was %d.\n", param, thePokemon.BaseExperience, dice)
	if dice > thePokemon.BaseExperience {
		c.Printer.Printf("You caught %s!\n", param)
		c.CaughtPokemons[param] = thePokemon
		c.Printer.Printf("There are now %d Pokémon in your Pokedex.\n", len(c.CaughtPokemons))
		// save caught pokemon for autocompletion
		internal.CurrentCompletionData["pokemon"] = append(internal.CurrentCompletionData["pokemon"], param)
	} else {
		c.Printer.Printf("%s escaped.\n", param)
	}
	return nil
}
