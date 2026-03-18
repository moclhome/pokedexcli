package pokeapi

import (
	"bootdev/go/pokedexcli/internal"
	"net/http"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      internal.Cache
}

// NewClient -
func NewClient(timeout time.Duration, flagDebug bool) Client {
	cache := internal.NewCache(10*time.Second, flagDebug)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache,
	}
}
