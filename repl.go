package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/flames31/Pok-dex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
	caughtPokemon   map[string]pokeapi.RespPokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}
		command := words[0]

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		val, ok := getCommands()[command]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := val.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	output := strings.Fields(text)
	return output
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},

		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"explore": {
			name:        "explore",
			description: "Explore a location",
			callback:    commandExplore,
		},

		"catch": {
			name:        "catch",
			description: "Catch a Pokemon!",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon you caught!",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show all the pokemon you caught!",
			callback:    commandPokedex,
		},
	}
}
