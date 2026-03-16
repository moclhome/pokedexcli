package pokeapi

import (
	"testing"
	"time"
)

func TestGetDataFromCacheOrInternet(t *testing.T) {
	pokeClient := NewClient(5*time.Second, true)

	cases := map[string]struct {
		url      string
		expected []byte
	}{
		"area_1": {
			url:      "https://pokeapi.co/api/v2/location-area?limit=1",
			expected: []byte("{\"count\":1104,\"next\":\"https://pokeapi.co/api/v2/location-area?offset=1&limit=1\",\"previous\":null,\"results\":[{\"name\":\"canalave-city-area\",\"url\":\"https://pokeapi.co/api/v2/location-area/1/\"}]}"),
		},
		"area_10": {
			url:      "https://pokeapi.co/api/v2/location-area?offset=9&limit=1",
			expected: []byte("{\"count\":1104,\"next\":\"https://pokeapi.co/api/v2/location-area?offset=10&limit=1\",\"previous\":\"https://pokeapi.co/api/v2/location-area?offset=8&limit=1\",\"results\":[{\"name\":\"fuego-ironworks-area\",\"url\":\"https://pokeapi.co/api/v2/location-area/10/\"}]}"),
		},
	}

	for testName, testData := range cases {
		result, err := pokeClient.GetDataFromCacheOrInternet(testData.url, true)
		if err != nil {
			t.Errorf("Test %s: Function returns error %v", testName, err)
		}
		if string(result) != string(testData.expected) {
			t.Errorf("Test %s:\nExpected:\n%s\nReturned:\n%s", testName, testData.expected, string(result))
		}
	}
}
