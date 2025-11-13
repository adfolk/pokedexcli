package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const PokeApiEndPoint string = "https://pokeapi.co/api/v2/location-area/"

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	conf := newConfig(PokeApiEndPoint)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		cmdName := words[0]
		cmd, exists := getCmds()[cmdName]
		if exists {
			cmd.callback(conf)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type Config struct {
	Next     string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *Config) error
}

func newConfig(ep string) *Config {
	c := Config{
		Next: ep,
	}
	return &c
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
			callback:    help,
		},
		"map": {
			name:        "map",
			description: "gets the next page of loc results",
			callback:    mapf,
		},
		"mapb": {
			name:        "mapb",
			description: "goes back to previous page of locations",
			callback:    mapb,
		},
	}
}
