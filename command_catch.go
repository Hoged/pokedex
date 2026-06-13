package main

import (
	"fmt"
	"errors"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemonResp, err := cfg.pokeapiClient.ListPokemon(name)
	if err != nil {
		return err
	}

	difficulty := pokemonResp.BaseXP / 4
	roll := rand.Intn(100) + 1

	fmt.Printf("Throwing a Pokeball at %v...\n", name)
	if roll < difficulty {
		fmt.Println(name, "escaped!")
		return nil
	}

	fmt.Println(name, "was caught!")

	cfg.caughtPokemon[name] = pokemonResp
	return nil
}

