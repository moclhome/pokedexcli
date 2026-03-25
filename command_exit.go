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
	c.Printer.Printf("Closing the Pokedex... Goodbye!\n")
	var jsonData pokeapi.UserData
	jsonData.Username = c.User
	for _, nextPokemon := range c.CaughtPokemons {
		jsonData.Pokedex = append(jsonData.Pokedex, nextPokemon)
	}
	data, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatal("Error during Marshal: ", err)
	}

	internal.WriteUserDataToFile(data)

	term.Restore(int(os.Stdin.Fd()), c.OldState)
	os.Exit(0)
	return nil
}
