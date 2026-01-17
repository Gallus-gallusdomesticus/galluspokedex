package main

import (
	"github.com/Gallus-gallusdomesticus/galluspokedex/internal/pokeapi"
)

type config struct {
	Next      *string
	Previous  *string
	ApiClient *pokeapi.Client
}
