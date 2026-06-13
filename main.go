package main

import (
	"time"

	"github.com/hoged/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.RespPokemon),
	}

	startRepl(cfg)
}
