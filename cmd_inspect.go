package main

import (
	"errors"
	"fmt"
)

func cmdInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Must enter a pokemon name")
	}
	printBasic := func(attr string, val any) {
		fmt.Printf("%s: %v\n", attr, val)
	}
	prettyPrint := func(key string, val any) {
		fmt.Printf(" -%s: %v\n", key, val)
	}
	if pokemon, ok := cfg.pokedex[args[0]]; ok {
		// do stuff
		printBasic("Name", pokemon.Name)
		printBasic("Height", pokemon.Height)
		printBasic("Weight", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			prettyPrint(stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, val := range pokemon.Types {
			fmt.Printf(" - %s\n", val.Type.Name)
		}
		return nil
	}
	return errors.New("You have not caught this pokemon yet")
}
