package main

import (
	"bootdev/go/pokedexcli/internal"
	"fmt"
	"os"

	"golang.org/x/term"
)

func startRepl(currentConfig *internal.Config) {
	internal.Registry = map[string]internal.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
			ParamType:   "",
		},
		"help": {
			Name:        "help",
			Description: "List all existing commands.",
			Callback:    commandHelp,
			ParamType:   "",
		},
		"map": {
			Name:        "map",
			Description: "List the next 20 location areas.",
			Callback:    commandMap,
			ParamType:   "",
		},
		"mapb": {
			Name:        "mapb",
			Description: "List the previous 20 location areas.",
			Callback:    commandMapb,
			ParamType:   "",
		},
		"explore": {
			Name:        "explore <area>",
			Description: "List the Pokémon located in an area.",
			Callback:    commandExplore,
			ParamType:   "area",
		},
		"catch": {
			Name:        "catch <pokemon>",
			Description: "Try to catch a Pokémon.",
			Callback:    commandCatch,
			ParamType:   "pokemon",
		},
		"inspect": {
			Name:        "inspect <pokemon>",
			Description: "Inspect a Pokémon that you have caught before.",
			Callback:    commandInspect,
			ParamType:   "pokemon",
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "List all the Pokémon that are in your Pokedex.",
			Callback:    commandPokedex,
		},
	}

	// Create a terminal using standard input/output
	terminal, err := createTerminal(currentConfig)
	if err != nil {
		currentConfig.Printer.Printf("%v", err)
	}

	// Configure autocompletion
	internal.CurrentCompletionData = make(map[string][]string)
	terminal.AutoCompleteCallback = internal.ContextAutocompletion
	for cmd := range internal.Registry {
		internal.CurrentCompletionData["command"] = append(internal.CurrentCompletionData["command"], cmd)
	}

	// Get user name for storing user dependent pokedex in a file
	user, err := terminal.ReadLine()
	if err != nil {
		currentConfig.Printer.Printf("Invalid input: %s", err)
		commandExit(currentConfig, "")
	}

	found, err := internal.FetchUserDataFromFile(currentConfig, user)
	if err != nil {
		currentConfig.Printer.Printf("Error fetching user data: %v", err)
	}

	defer func() {
		term.Restore(int(os.Stdin.Fd()), currentConfig.OldState)
		currentConfig.Printer.Println("defer in startRepl - restored old status")
		if err := recover(); err != nil {
			currentConfig.Printer.Printf("Fatal Error '%v'! But recovered.", err)
		}
	}()

	back := ""
	if found {
		back = "back "
	}

	currentConfig.User = user
	currentConfig.Printer.Printf("\nWelcome %sto the Pokedex, %s!\nTo display a list of available commands, use \"help\".\n", back, user)
	if found {
		currentConfig.Printer.Printf("You already have %d Pokémon in your Pokedex.\n", len(currentConfig.CaughtPokemons))
	}

	terminal.SetPrompt("Pokedex >")
	// Read lines interactively
	for {
		line, err := terminal.ReadLine()
		if err != nil {
			currentConfig.Printer.Printf("Invalid input: %s", err)
			break
		}

		cleanedInput := internal.CleanInput(line)

		theCommand := cleanedInput[0]
		var firstParam string
		if len(cleanedInput) > 1 {
			firstParam = cleanedInput[1]
		}

		commandStruct, ok := internal.Registry[theCommand]
		if !ok {
			currentConfig.Printer.Println("Unknown command\n")
		} else {
			err := commandStruct.Callback(currentConfig, firstParam)
			if err != nil {
				currentConfig.Printer.Printf("Error: %v", err)
			}
		}
	}
}

func createTerminal(c *internal.Config) (*term.Terminal, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	c.OldState = oldState
	if err != nil {
		return nil, fmt.Errorf("Error putting the terminal to raw mode: %v\n", err)
	}
	terminal := term.NewTerminal(os.Stdin, "Please type your user name (will be part of a filename):")

	wrapper := internal.TerminalWrapper{T: terminal}
	c.Printer = wrapper
	return terminal, nil
}
