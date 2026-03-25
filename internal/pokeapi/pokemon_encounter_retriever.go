package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetPokemonEncounter(areaName string, flagDebug bool) (areaWithEncounter, error) {
	url := baseURL + "/location-area/" + areaName + "/"

	data, err := c.GetDataFromCacheOrInternet(url, flagDebug)
	if err != nil {
		return areaWithEncounter{}, err
	}

	var certainArea areaWithEncounter
	if err = json.Unmarshal(data, &certainArea); err != nil {
		return areaWithEncounter{}, err
	}

	return certainArea, nil
}
