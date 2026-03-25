package pokeapi

import (
	"encoding/json"
	"log"
)

func (c *Client) GetLocationAreas(url string, flagDebug bool) (locationArea, error) {
	firstUrl := baseURL + "/location-area"
	if url == "" || url == "start" { // "" should not happen, but to be on the save side
		url = firstUrl
	}

	if flagDebug {
		log.Printf("Provide locations for url %s\n", url)
	}
	data, err := c.GetDataFromCacheOrInternet(url, flagDebug)
	if err != nil {
		return locationArea{}, err
	}

	var locAreas locationArea
	if json.Unmarshal(data, &locAreas) != nil {
		return locationArea{}, err
	}

	return locAreas, nil
}
