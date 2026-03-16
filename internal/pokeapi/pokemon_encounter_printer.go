package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) PrintPokemonEncounters(areaName string, flagDebug bool) error {
	url := baseURL + "/location-area/" + areaName + "/"

	data, err := c.GetDataFromCacheOrInternet(url, flagDebug)
	if err != nil {
		return err
	}

	fmt.Println("Pokemon found:")

	var certainArea areaWithEncounter
	if err = json.Unmarshal(data, &certainArea); err != nil {
		return err
	}

	for _, encounter := range certainArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
