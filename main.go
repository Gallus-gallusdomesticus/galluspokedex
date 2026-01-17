package main

import "github.com/Gallus-gallusdomesticus/galluspokedex/internal/pokeapi"

func main() {

	cfg := &config{
		ApiClient: pokeapi.NewClient(nil),
	}
	startRepl(cfg)
}
