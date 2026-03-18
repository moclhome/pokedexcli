package pokeapi

import (
	"bootdev/go/pokedexcli/internal"
	"encoding/json"
)

func (c *Client) PrintPokemonEncounters(areaName string, flagDebug bool, p internal.Printer) error {
	url := baseURL + "/location-area/" + areaName + "/"

	data, err := c.GetDataFromCacheOrInternet(url, flagDebug)
	if err != nil {
		return err
	}

	p.Println("Pokemon found:")

	var certainArea areaWithEncounter
	if err = json.Unmarshal(data, &certainArea); err != nil {
		return err
	}

	for _, encounter := range certainArea.PokemonEncounters {
		p.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
