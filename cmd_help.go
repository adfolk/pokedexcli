package main

import (
	"fmt"
)

func cmdHelp(cfg *config, args ...string) error {
	// TODO: dynamically generate usage section by iterating the commands map
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for key, value := range getCmds() {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}
