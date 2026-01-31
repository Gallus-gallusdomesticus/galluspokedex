package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: explore <location-name>")
	}

	areaName := args[0]

	loc, err := cfg.ApiClient.GetLocation(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")
	for _, loca := range loc.PokemonEncounters {
		fmt.Println(loca.Pokemon.Name)
	}

	return nil

}
