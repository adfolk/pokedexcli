package main

import (
	"errors"
	"fmt"
)

func cmdExplore(cfg *config, args ...string) error {
	switch numArgs := len(args); numArgs {
	case 0:
		return errors.New("you must provide a location name")
	case 1:
		// success
		location, err := cfg.pokeapiClient.GetLocation(args[0])
		if err != nil {
			return err
		}
		fmt.Printf("Exploring %s...\n", location.Name)
		fmt.Println("Found Pokemon: ")
		for _, enc := range location.PokemonEncounters {
			fmt.Printf(" - %s\n", enc.Pokemon.Name)
		}
	default:
		return errors.New("Too many arguments")
	}

	return nil
}
