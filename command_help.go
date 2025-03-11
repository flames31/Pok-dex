package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, v := range getCommands() {
		fmt.Printf("\n%v: %v\n", v.name, v.description)
	}
	return nil
}
