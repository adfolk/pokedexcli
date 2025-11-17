package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func cmdCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Must enter a pokemon name")
	}
	target, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", target.Name)
	baseXP := target.BaseExperience
	minXP := 25
	startingProb := 0.9
	floor := 0.1
	m := -0.0005
	prob := startingProb + m*(float64(baseXP-minXP))
	if prob < floor {
		prob = floor
	}
	attempt := rand.Float64()

	if attempt < prob {
		fmt.Printf("Success! Captured %s\n", target.Name)
		cfg.pokedex[target.Name] = target
		return nil
	}
	fmt.Printf("Failed to catch %s\n", target.Name)
	return nil
}
