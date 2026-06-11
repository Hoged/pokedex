package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"github.com/hoged/pokedex/internal/pokeapi"
)

type cliCommands struct {
	name		string
	description	string
	callback	func(*config) error
}

type config struct {
	pokeapiClient   		pokeapi.Client
	nextLocationsURL		*string
	prevLocationsURL		*string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		userInput := cleanInput(scanner.Text())
		if len(userInput) == 0 {
			continue
		}

		command := userInput[0]

		com, exists := getCommands()[command]
		if exists {
			err := com.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
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
		"map": {
			name:			"map",
			description:	"Displays names of next 20 location areas in Pokemon world",
			callback:		commandMap,
		},
		"mapb": {
			name:			"mapb",
			description:	"Displays names of previous 20 location areas in Pokemon world",
			callback:		commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	cleanText := strings.Fields(strings.ToLower(text))
	return cleanText
}