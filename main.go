package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	newscan := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		newscan.Scan()
		input := newscan.Text()
		cleaned := cleanInput(input)
		fmt.Println("Your command was:", cleaned[0])
	}
}
