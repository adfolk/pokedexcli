package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type config struct {
	Previous string `json:"previous"`
	Next     string `json:"next"`
}

func updateConfig(oldConf *config, raw []byte) error {
	var newConf config
	err := json.Unmarshal(raw, &newConf)
	if err != nil {
		fmt.Println(err)
	}
	*oldConf = newConf
	return nil
}

func startRepl() {
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
			c := command.confPtr
			err := command.callback(c)
			if err != nil {
				fmt.Println(err)
			}
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
	callback    func(c *config) error
	confPtr     *config
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
			confPtr:     nil,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			confPtr:     nil,
		},
		"map": {
			name:        "map",
			description: "Gets a page of locations",
			callback:    getLocations,
			confPtr:     &mapConfig,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets and displays the previous page of locations",
			callback:    getPrevLocations,
			confPtr:     &mapConfig,
		},
	}
}
