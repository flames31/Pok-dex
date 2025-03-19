package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {

	if len(args) < 1 {
		return fmt.Errorf("need a location you bum")
	}
	location := args[0]

	resp, err := cfg.pokeapiClient.ExploreLocation(location)

	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", location)
	for _, pok_encounter := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pok_encounter.Pokemon.Name)
	}

	return nil
}
