package pokeapi

import (
	"bootdev/go/pokedexcli/internal"
	"encoding/json"
)

func (c *Client) PrintLocationAreas(url string, flagDebug bool, p internal.Printer) (locationArea, error) {
	firstUrl := baseURL + "/location-area"
	if url == "" || url == "start" { // "" should not happen, but to be on the save side
		url = firstUrl
	}

	if flagDebug {
		p.Printf("Provide locations for url %s\n", url)
	}
	data, err := c.GetDataFromCacheOrInternet(url, flagDebug)
	if err != nil {
		return locationArea{}, err
	}

	var locAreas locationArea
	if json.Unmarshal(data, &locAreas) != nil {
		return locationArea{}, err
	}

	for _, locationArea := range locAreas.Results {
		p.Printf("%s\n", locationArea.Name)
	}
	return locAreas, nil
}
