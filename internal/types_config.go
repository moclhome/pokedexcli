package internal

import (
	"bootdev/go/pokedexcli/internal/pokeapi"

	"golang.org/x/term"
)

type Config struct {
	PokeapiClient  pokeapi.Client
	Next           string
	Previous       string
	FlagDebug      bool
	CaughtPokemons map[string]pokeapi.Pokemon
	OldState       *term.State
	Printer        Printer
	User           string
}
