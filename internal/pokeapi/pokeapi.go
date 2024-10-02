package pokeapi

import (
	"jwald3/pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInverval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInverval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
