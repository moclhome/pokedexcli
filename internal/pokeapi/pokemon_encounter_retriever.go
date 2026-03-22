package pokeapi

import (
	"bootdev/go/pokedexcli/internal"
	"encoding/json"
)

func (c *Client) GetPokemonEncounter(areaName string, flagDebug bool, p internal.Printer) (areaWithEncounter, error) {
	url := baseURL + "/location-area/" + areaName + "/"

	data, err := c.GetDataFromCacheOrInternet(url, flagDebug)
	if err != nil {
		return areaWithEncounter{}, err
	}

	p.Println("Pokemon found:")

	var certainArea areaWithEncounter
	if err = json.Unmarshal(data, &certainArea); err != nil {
		return areaWithEncounter{}, err
	}

	return certainArea, nil
}
