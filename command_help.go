package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	commands := getCommands()

	for _, k := range commands { //iterate over the map
		cmdname := k.name
		cmddesc := k.description
		fmt.Printf("%s: %s\n", cmdname, cmddesc)

	}

	return nil
}
