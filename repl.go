package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display 20 location name",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 location name",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <area>",
			description: "List the pokemon on the specified area",
			callback:    commandExplore,
		},
	}
}

func startRepl(cfg *config) {
	newscan := bufio.NewScanner(os.Stdin) //for the input

	for {
		fmt.Print("Pokedex > ")      //print the prompt
		newscan.Scan()               //scan what being input
		input := newscan.Text()      //convert it to string
		cleaned := cleanInput(input) //clean the string

		if len(cleaned) == 0 { //conditional to prevent panic if no input given
			continue
		}

		args := []string{} //for arguments calling
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		commandName := cleaned[0] //the first word is the command
		commands := getCommands() //access the map

		cmd, ok := commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)           //lower all uppercase
	nowhitespace := strings.TrimSpace(lowered) //remove all trailing whitespace

	separated := strings.Fields(nowhitespace) //separate into slice

	return separated
}
