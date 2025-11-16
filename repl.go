package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/adfolk/pokedexcli/internal/pokeapi"
	"github.com/adfolk/pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient    pokeapi.Client
	pokeCache        *pokecache.Cache
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		cmd, exists := getCmds()[cmdName]
		if exists {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func getCmds() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    cmdExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    cmdHelp,
		},
		"catch": {
			name:        "catch",
			description: "catch a pokemon",
			callback:    cmdCatch,
		},
		"map": {
			name:        "map",
			description: "gets the next page of loc results",
			callback:    cmdMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "goes back to previous page of locations",
			callback:    cmdMapb,
		},
		"explore": {
			name:        "explore",
			description: "list all pokemon in an area",
			callback:    cmdExplore,
		},
	}
}
