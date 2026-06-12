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
	callback	func(*config, ...string) error
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
		args := []string{}

		if len(userInput) > 1 {
			args = userInput[1:]
		}

		com, exists := getCommands()[command]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := com.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
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
		"explore": {
			name:			"explore <location_name>",
			description:	"Displays names of Pokemons you can find in a given area",
			callback:		commandExplore,
		},
	}
}

func cleanInput(text string) []string {
	cleanText := strings.Fields(strings.ToLower(text))
	return cleanText
}