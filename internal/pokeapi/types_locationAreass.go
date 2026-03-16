package pokeapi

type locationArea struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
}

type areaWithEncounter struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
