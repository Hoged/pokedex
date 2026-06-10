package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type cliCommands struct {
	name		string
	description	string
	callback	func() error
}

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

		com, ok := getCommands()[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		com.callback()
	}
}

func getCommands() map[string]cliCommands {
	return map[string]cliCommands{
		"exit": {
			name:			"exit",
			description:	"Exit the Pokedex",
			callback:		commandExit, 
		},
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback:		commandHelp,
		},
	}
}

func cleanInput(text string) []string {
	cleanText := strings.Fields(strings.ToLower(text))
	return cleanText
}