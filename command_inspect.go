package main

import (
	"bootdev/go/pokedexcli/internal"
	"strings"
)

func commandInspect(c *internal.Config, param string) error {
	movesPerLine := 8
	if pokemon, found := c.CaughtPokemons[param]; !found {
		c.Printer.Printf("You have not yet caught a Pokémon named %s.\n", param)
	} else {
		c.Printer.Printf("Name: %s\n", pokemon.Name)
		c.Printer.Printf("Height: %d\n", pokemon.Height)
		c.Printer.Printf("Weight: %d\n", pokemon.Weight)
		c.Printer.Printf("Base Experience: %d\n", pokemon.BaseExperience)
		c.Printer.Println("Stats:")
		for _, stat := range pokemon.Stats {
			c.Printer.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		c.Printer.Println("Abilities:")
		for i := 0; i < len(pokemon.Abilities); i++ {
			c.Printer.Printf("  - %s\n", pokemon.Abilities[i].Ability.Name)
		}
		for i := 0; i < len(pokemon.Types); i++ { // why is this not working with range?
			c.Printer.Printf("  - %s\n", pokemon.Types[i].Type.Name)
		}
		c.Printer.Printf("Moves (%d):\n", len(pokemon.Moves))
		for i := 0; i < len(pokemon.Moves)/movesPerLine; i++ {
			var moves10 []string
			for j := i * movesPerLine; j < (i+1)*movesPerLine; j++ {
				moves10 = append(moves10, pokemon.Moves[j].Move.Name)
			}
			moves10string := strings.Join(moves10, ", ")
			c.Printer.Println("   " + moves10string + ",")
		}
		var movesRest []string
		for j := 0; j < len(pokemon.Moves)%movesPerLine; j++ {
			index := int(len(pokemon.Moves)/movesPerLine)*movesPerLine + j
			movesRest = append(movesRest, pokemon.Moves[index].Move.Name)
		}
		movesRestString := strings.Join(movesRest, ", ")
		c.Printer.Println("   " + movesRestString)
		c.Printer.Println("Types:")
		/*for _, type := range pokemon.Types { // I get the error "expected 1 expression" TODO
			fmt.Printf("  - %s\n", type.Type.Name)
		}*/
	}
	return nil
}
