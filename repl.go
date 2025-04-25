package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var supportedCommands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the pokedex",
		callback:    commandExit,
	},
}

func init() {
	supportedCommands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
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
		if err := runCommand(commandName, supportedCommands); err != nil {
			fmt.Printf("%v", err)
		}
	}
}

func cleanInput(text string) []string {
	results := strings.Fields(strings.ToLower(text))
	return results
}

func runCommand(text string, registry map[string]cliCommand) error {
	cmd, exists := registry[text]
	if exists {
		err := cmd.callback()
		if err != nil {
			return err
		}
		return nil // execution should have succeeded
	} else {
		return fmt.Errorf("Unknown command\n")
	}
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range supportedCommands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	return nil
}
