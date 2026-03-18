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
	callback    func(*config, string) error
}

var registry map[string]cliCommand

// Usage:
/*
terminal := term.NewTerminal(os.Stdin, "> ")
wrapper := &TerminalWrapper{Terminal: terminal}
wrapper.Printf("Value: %d, Name: %s\n", value, name)
*/

func startRepl(currentConfig *config) {
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

	fmt.Println("Welcome to the Pokedex!\nTo display a list of available commands, use \"help\".")

	// Create a terminal using standard input/output
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	currentConfig.oldState = oldState
	if err != nil {
		currentConfig.printer.Printf("Error putting the terminal to raw mode: %v\n", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	terminal := term.NewTerminal(os.Stdin, "Pokedex >")
	wrapper := internal.TerminalWrapper{T: terminal}
	currentConfig.printer = wrapper

	// Read lines interactively
	for {
		line, err := terminal.ReadLine()
		if err != nil {
			currentConfig.printer.Printf("Invalid input: %s", err)
			break
		}
		// Process line

		/*scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("Pokedex >")
			scanner.Scan()
			input := scanner.Text()
			if err := scanner.Err(); err != nil {
				fmt.Printf("Invalid input: %s", err)
			}*/
		cleanedInput := CleanInput(line)

		theCommand := cleanedInput[0]
		var firstParam string
		if len(cleanedInput) > 1 {
			firstParam = cleanedInput[1]
		}
		if theCommand == "test" {
			for i := 0; i < 5; i++ {
				currentConfig.printer.Printf("test %d Printf\n", i)
			}
		}

		commandStruct, ok := registry[theCommand]
		if !ok {
			currentConfig.printer.Println("Unknown command\n")
		} else {
			err := commandStruct.callback(currentConfig, firstParam)
			if err != nil {
				currentConfig.printer.Printf("Error: %v", err)
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
