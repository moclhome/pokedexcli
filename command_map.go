package main

func commandMap(c *config, param string) error {
	if c.Next == "" {
		c.printer.Println("you're on the last page")
		return nil
	}

	locAreas, err := c.pokeapiClient.PrintLocationAreas(c.Next, c.flagDebug, c.printer)
	if err != nil {
		return err
	}

	c.Next = locAreas.Next
	c.Previous = locAreas.Previous
	return nil
}

func commandMapb(c *config, param string) error {
	if c.Previous == "" {
		c.printer.Println("you're on the first page")
		return nil
	}

	locAreas, err := c.pokeapiClient.PrintLocationAreas(c.Previous, c.flagDebug, c.printer)
	if err != nil {
		return err
	}

	c.Next = locAreas.Next
	c.Previous = locAreas.Previous
	return nil
}
