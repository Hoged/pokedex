package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		userInput := cleanInput(scanner.Text())
		if len(userInput) == 0 {
			continue
		}

		command := userInput[0]

		fmt.Printf("Your command was: %v\n", command)
	}
}

func cleanInput(text string) []string {
	cleanText := strings.Fields(strings.ToLower(text))
	return cleanText
}