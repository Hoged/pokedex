package main

import (
	"fmt"
	"errors"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	areaResp, err := cfg.pokeapiClient.ListArea(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", name)
	fmt.Println("Found pokemon:")
	for _, encounter := range areaResp.PokemonEncounters {
		fmt.Println("-", encounter.Pokemon.Name)
	}
	return nil
}

