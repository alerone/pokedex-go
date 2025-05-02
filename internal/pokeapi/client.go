package pokeapi

import (
	"net/http"
	"time"

	"github.com/alerone/pokedex-go/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
	pokedex    map[string]Pokemon
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
		pokedex: make(map[string]Pokemon),
	}
}
