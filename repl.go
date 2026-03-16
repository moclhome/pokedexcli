package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

var registry map[string]cliCommand

func startRepl(currentConfig config) {
	registry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "List all existing commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "List the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <area>",
			description: "List the Pokémon located in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Try to catch a Pokémon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Inspect a Pokémon that you have caught before",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all the Pokémon that are in your Pokedex",
			callback:    commandPokedex,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Printf("Invalid input: %s", err)
		}
		cleanedInput := CleanInput(input)

		theCommand := cleanedInput[0]
		var firstParam string
		if len(cleanedInput) > 1 {
			firstParam = cleanedInput[1]
		}
		commandStruct, ok := registry[theCommand]
		if !ok {
			fmt.Printf("Unknown command\n")
		} else {
			err := commandStruct.callback(&currentConfig, firstParam)
			if err != nil {
				fmt.Println("Error: ", err)
			}
		}
	}
}

/**
 *
 */
func CleanInput(text string) []string {
	var returnSlice []string
	returnSlice = strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
	return returnSlice
}
