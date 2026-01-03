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
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl() {
	newscan := bufio.NewScanner(os.Stdin) //for the input

	for {
		fmt.Print("Pokedex > ")      //print the prompt
		newscan.Scan()               //scan what being input
		input := newscan.Text()      //convert it to string
		cleaned := cleanInput(input) //clean the string

		if len(cleaned) == 0 { //conditional to prevent panic if no input given
			continue
		}

		commandName := cleaned[0] //the first word is the command
		commands := getCommands() //access the map

		cmd, ok := commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			return
		}

		err := cmd.callback()
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
