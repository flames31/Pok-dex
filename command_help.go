package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, v := range getCommands() {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	return nil
}
