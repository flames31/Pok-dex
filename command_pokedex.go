package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You have not caught any pokemon loser")
	}

	for pok := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pok)
	}

	return nil
}
