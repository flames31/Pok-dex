package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("need a pokemon name bruv")
	}
	pokemonName := args[0]

	if _, ok := cfg.caughtPokemon[pokemonName]; !ok {
		fmt.Println("You have not caught this pokemon foo. Cant inspect!")
		return nil
	}

	pokemonInfo := cfg.caughtPokemon[pokemonName]

	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pType := range pokemonInfo.Types {
		fmt.Printf("  - %s\n", pType.Type.Name)
	}

	return nil
}
