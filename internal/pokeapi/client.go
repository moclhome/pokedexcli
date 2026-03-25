package pokeapi

import (
	"bootdev/go/pokedexcli/internal/pokecach"
	"net/http"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      pokecach.Cache
}

// NewClient -
func NewClient(timeout time.Duration, flagDebug bool) Client {
	cache := pokecach.NewCache(10*time.Second, flagDebug)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache,
	}
}
