package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: catch <pokemon-name>")
	}

	pokemonName := args[0]

	pok, err := cfg.ApiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	source := rand.NewSource(time.Now().UnixNano()) //create seed for randomness
	r := rand.New(source)
	throwNum := r.Intn(701)

	fmt.Printf("Throwing a Pokeball at %s...\n", pok.Name)
	if throwNum >= pok.Exp {

		fmt.Printf("%s was caught!\n", pok.Name)
		if _, exists := cfg.pokedex[pok.Name]; !exists {
			cfg.pokedex[pok.Name] = pok
			fmt.Println("You may now inspect it with the inspect command.")
		}

	} else {
		fmt.Printf("%s escaped!\n", pok.Name)
	}

	return nil

}
