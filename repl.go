package main

import (
	"bootdev/go/pokedexcli/internal"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*internal.Config, string) error
	paramType   string
}

var registry map[string]cliCommand

func startRepl(currentConfig *internal.Config) {
	registry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
			paramType:   "",
		},
		"help": {
			name:        "help",
			description: "List all existing commands.",
			callback:    commandHelp,
			paramType:   "",
		},
		"map": {
			name:        "map",
			description: "List the next 20 location areas.",
			callback:    commandMap,
			paramType:   "",
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous 20 location areas.",
			callback:    commandMapb,
			paramType:   "",
		},
		"explore": {
			name:        "explore <area>",
			description: "List the Pokémon located in an area.",
			callback:    commandExplore,
			paramType:   "area",
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Try to catch a Pokémon.",
			callback:    commandCatch,
			paramType:   "pokemon",
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Inspect a Pokémon that you have caught before.",
			callback:    commandInspect,
			paramType:   "pokemon",
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all the Pokémon that are in your Pokedex.",
			callback:    commandPokedex,
		},
	}

	// Create a terminal using standard input/output
	terminal, err := createTerminal(currentConfig)
	if err != nil {
		currentConfig.Printer.Printf("%v", err)
	}

	// Configure autocompletion
	CurrentCompletionData = make(map[string][]string)
	terminal.AutoCompleteCallback = ContextAutocompletion
	for cmd := range registry {
		CurrentCompletionData["command"] = append(CurrentCompletionData["command"], cmd)
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
		currentConfig.File.Close()
	}()

	back := ""
	if found {
		back = "back "
	}

	currentConfig.User = user
	currentConfig.Printer.Printf("Welcome %sto the Pokedex, %s!\nTo display a list of available commands, use \"help\".\n", back, user)
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

		cleanedInput := CleanInput(line)

		theCommand := cleanedInput[0]
		var firstParam string
		if len(cleanedInput) > 1 {
			firstParam = cleanedInput[1]
		}

		commandStruct, ok := registry[theCommand]
		if !ok {
			currentConfig.Printer.Println("Unknown command\n")
		} else {
			err := commandStruct.callback(currentConfig, firstParam)
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
	terminal := term.NewTerminal(os.Stdin, "Please type your user name:")

	wrapper := internal.TerminalWrapper{T: terminal}
	c.Printer = wrapper
	return terminal, nil
}

/**
 *
 */
func CleanInput(text string) []string {
	var returnSlice []string
	returnSlice = strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
	return returnSlice
}
