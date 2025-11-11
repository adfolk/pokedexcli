package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cmdExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func help() error {
	// TODO: dynamically generate usage section by iterating the commands map
	fmt.Println(`Welcome to the Pokedex!
		Usage:

		help: Displays a help message
		exit: Exit the Pokedex`)
	return nil
}

var commands map[string]cliCommand = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    cmdExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    help,
	},
}
