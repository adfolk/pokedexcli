package main

import "fmt"

func cmdPokedex(cfg *config, args ...string) error {
	pretty := func(name string) {
		fmt.Printf(" - %s\n", name)
	}
	fmt.Println("Your Pokedex:")
	for _, val := range cfg.pokedex {
		pretty(val.Name)
	}
	return nil
}
