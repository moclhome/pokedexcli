package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func commandCatch(c *config, param string) error {
	if _, found := c.caughtPokemons[param]; found {
		fmt.Printf("You already have %s in your Pokédex. Try someone else!\n", param)
		return nil
	}
	thePokemon, err := c.pokeapiClient.GetPokemon(param, c.flagDebug)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			fmt.Printf("There is no Pokémon named %s.\n", param)
			return nil
		}
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", param)
	dice := rand.Intn(250)
	fmt.Printf("%s has experience %d. Throwing the dice... It was %d.\n", param, thePokemon.BaseExperience, dice)
	if dice > thePokemon.BaseExperience {
		fmt.Printf("You caught %s!\n", param)
		c.caughtPokemons[param] = thePokemon
		fmt.Printf("There are now %d Pokémon in your Pokedex.\n", len(c.caughtPokemons))
	} else {
		fmt.Printf("%s escaped.\n", param)
	}
	return nil
}
