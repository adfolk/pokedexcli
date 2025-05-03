package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/adfolk/bootdev-guided/pokedexcli/internal/pokeapi"
	"github.com/adfolk/bootdev-guided/pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient    pokeapi.Client
	pokecacheCache   *pokecache.Cache
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		if err := s.Err(); err != nil {
			resErr := fmt.Errorf("Error scanning user input: %w", err)
			fmt.Printf("%v", resErr)
		}

		words := cleanInput(s.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	results := strings.Fields(strings.ToLower(text))
	return results
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Gets a page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets and displays the previous page of locations",
			callback:    commandMapb,
		},
	}
}
