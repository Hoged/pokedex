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
	callback	func(*config) error
}

type jsonMap struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type config struct {
	next		string
	previous	string
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := config{
		next:		"https://pokeapi.co/api/v2/location-area/",
		previous:	"",
	}

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
		com.callback(&conf)
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
			callback:		commandMapBack,
		},
	}
}

func cleanInput(text string) []string {
	cleanText := strings.Fields(strings.ToLower(text))
	return cleanText
}