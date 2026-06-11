package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n\n")

	for _, cmd := range getCommands() {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}