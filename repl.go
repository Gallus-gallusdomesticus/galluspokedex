package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

		fmt.Printf("Your command was: %s\n", cleaned[0])
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)           //lower all uppercase
	nowhitespace := strings.TrimSpace(lowered) //remove all trailing whitespace

	separated := strings.Fields(nowhitespace) //separate into slice

	return separated
}
