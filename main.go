package main

import (
	"time"

	"bootdev/go/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient  pokeapi.Client
	Next           string
	Previous       string
	flagDebug      bool
	caughtPokemons map[string]pokeapi.Pokemon
}

func main() {
	var flagDebug bool = false
	pokeClient := pokeapi.NewClient(20*time.Second, flagDebug)
	currentConfig := config{
		pokeapiClient:  pokeClient,
		Next:           "start",
		Previous:       "",
		flagDebug:      flagDebug,
		caughtPokemons: make(map[string]pokeapi.Pokemon),
	}

	startRepl(currentConfig)
}
