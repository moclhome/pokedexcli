package main

import (
	"os"

	"golang.org/x/term"
)

func commandExit(c *config, param string) error {
	c.printer.Printf("Closing the Pokedex... Goodbye!\n")
	term.Restore(int(os.Stdin.Fd()), c.oldState)
	os.Exit(0)
	return nil
}
