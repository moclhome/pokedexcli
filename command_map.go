package main

import "bootdev/go/pokedexcli/internal"

func commandMap(c *internal.Config, param string) error {
	if c.Next == "" {
		c.Printer.Println("you're on the last page")
		return nil
	}

	locAreas, err := c.PokeapiClient.GetLocationAreas(c.Next, c.FlagDebug)
	if err != nil {
		return err
	}

	c.Next = locAreas.Next
	c.Previous = locAreas.Previous
	for _, area := range locAreas.Results {
		c.Printer.Printf("%s\n", area.Name)
		// save area for autocompletion
		CurrentCompletionData["area"] = append(CurrentCompletionData["area"], area.Name)
	}
	return nil
}

func commandMapb(c *internal.Config, param string) error {
	if c.Previous == "" {
		c.Printer.Println("you're on the first page")
		return nil
	}

	locAreas, err := c.PokeapiClient.GetLocationAreas(c.Previous, c.FlagDebug)
	if err != nil {
		return err
	}

	c.Next = locAreas.Next
	c.Previous = locAreas.Previous
	for _, area := range locAreas.Results {
		c.Printer.Printf("%s\n", area.Name)
		// save area for autocompletion
		CurrentCompletionData["area"] = append(CurrentCompletionData["area"], area.Name)
	}
	return nil
}
