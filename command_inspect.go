package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: inspect <pokemon-name>")
	}

	pokemonName := args[0]
	pokeList, exists := cfg.pokedex[pokemonName]
	if exists {
		fmt.Printf("Name: %s\n", pokeList.Name)
		fmt.Printf("Height: %d\n", pokeList.Height)
		fmt.Printf("Weight: %d\n", pokeList.Weight)

		fmt.Println("Stats: ")
		for _, stat := range pokeList.Stats {
			fmt.Printf("	-%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types: ")
		for _, typelist := range pokeList.Types {
			fmt.Println("	-", typelist.Type.Name)
		}

	} else {
		fmt.Printf("%s is not yet caught!\n", pokemonName)
	}
	return nil
}
