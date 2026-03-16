package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetPokemon(pokemonName string, flagDebug bool) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName + "/"

	data, err := c.GetDataFromCacheOrInternet(url, flagDebug)
	if err != nil {
		return Pokemon{}, err
	}
	var thePokemon Pokemon
	if err = json.Unmarshal(data, &thePokemon); err != nil {
		return Pokemon{}, err
	}
	return thePokemon, nil
}
