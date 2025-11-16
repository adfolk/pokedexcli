package main

import (
	"errors"
	"fmt"
)

func cmdCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Must enter a pokemon name")
	}
	target := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", target)

	// Get info about the specific pokemon by name

	// Generate catching chance with math/rand and pokemon's base experience

	// If caught, add pokemon to user's pokedex

	return nil
}
