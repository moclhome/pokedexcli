package main

import (
	"bootdev/go/pokedexcli/internal"
	"bootdev/go/pokedexcli/internal/pokeapi"
	"encoding/json"
	"log"
	"os"

	"golang.org/x/term"
)

func commandExit(c *internal.Config, param string) error {
	c.Printer.Printf("Closing the Pokedex... Goodbye, %s!\n", c.User)
	var jsonData []pokeapi.Pokemon
	for _, nextPokemon := range c.CaughtPokemons {
		jsonData = append(jsonData, nextPokemon)
	}
	data, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatal("Error during Marshal: ", err)
	}

	internal.WriteUserDataToFile(data, c.User)

	term.Restore(int(os.Stdin.Fd()), c.OldState)
	os.Exit(0)
	return nil
}
