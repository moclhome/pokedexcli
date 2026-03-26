package main

import (
	"bootdev/go/pokedexcli/internal"
	"bootdev/go/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	var flagDebug bool = false
	pokeClient := pokeapi.NewClient(20*time.Second, flagDebug)
	currentConfig := internal.Config{
		PokeapiClient:  pokeClient,
		Next:           "start",
		Previous:       "",
		FlagDebug:      flagDebug,
		CaughtPokemons: make(map[string]pokeapi.Pokemon),
		Printer:        nil,
	}

	startRepl(&currentConfig)
	commandExit(&currentConfig, "")
}
