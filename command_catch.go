package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("need a pokemon name bruv")
	}
	pokemonName := args[0]

	if _, ok := cfg.caughtPokemon[pokemonName]; ok {
		fmt.Println("You already caught this pokemon foo")
		return nil
	}

	pokemonInfo, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	pokExp := pokemonInfo.BaseExperience

	maxExperience := 340
	minProbability := 0.05
	maxProbability := 0.95

	probability := maxProbability - (float64(pokExp)/float64(maxExperience))*(maxProbability-minProbability)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	for i := 0; i < 3; i++ {
		fmt.Println(".")
		time.Sleep(1 * time.Second)
	}
	if rand.Float64() < probability {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.caughtPokemon[pokemonName] = pokemonInfo
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
