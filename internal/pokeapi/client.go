package pokeapi

import (
	"net/http"
	"time"

	"github.com/Gallus-gallusdomesticus/galluspokedex/internal/pokecache"
)

type Client struct {
	httpClient *http.Client
	cache      *pokecache.Cache
}

//New Client for PokeApi

func NewClient(httpClient *http.Client, cacheInterval time.Duration) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	return &Client{
		httpClient: httpClient,
		cache:      pokecache.NewCache(cacheInterval),
	}
}
