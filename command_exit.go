package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing Pokedex... bye bitch!")
	os.Exit(0)
	return nil
}
