package main

import (
	"net/http"
	"time"

	"github.com/Gallus-gallusdomesticus/galluspokedex/internal/pokeapi"
)

func main() {

	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	pokeClient := pokeapi.NewClient(httpClient, 5*time.Minute)
	cfg := &config{
		ApiClient: pokeClient,
	}
	startRepl(cfg)
}
