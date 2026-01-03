package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replbase() {
	newscan := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		newscan.Scan()
		input := newscan.Text()
		cleaned := cleanInput(input)
		fmt.Println("Your command was:", cleaned[0])
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)           //lower all uppercase
	nowhitespace := strings.TrimSpace(lowered) //remove all trailing whitespace

	separated := strings.Fields(nowhitespace) //separate into slice

	return separated
}
