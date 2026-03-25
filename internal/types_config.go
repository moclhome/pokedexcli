package internal

import (
	"bootdev/go/pokedexcli/internal/pokeapi"
	"os"

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
	File           *os.File
	UserData       pokeapi.UserData
}
