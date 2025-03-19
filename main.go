package main

import (
	"time"

	"github.com/flames31/Pok-dex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.RespPokemon{},
	}
	startRepl(cfg)
}
