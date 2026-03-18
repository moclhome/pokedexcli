package main

import (
	"time"

	"bootdev/go/pokedexcli/internal"
	"bootdev/go/pokedexcli/internal/pokeapi"

	"golang.org/x/term"
)

type config struct {
	pokeapiClient  pokeapi.Client
	Next           string
	Previous       string
	flagDebug      bool
	caughtPokemons map[string]pokeapi.Pokemon
	oldState       *term.State
	printer        internal.Printer
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
		printer:        nil,
	}

	startRepl(&currentConfig)
}
