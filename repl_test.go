package main

import (
	"fmt"
	"testing"
)

var cfg = &config{}

func nerfedCommandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	return nil
}

func getNerfedCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exitNerfed",
			description: "Pretend to exit the pokedex",
			callback:    nerfedCommandExit,
		},
	}
}

func startTestRepl(cfg *config, str string) (cmdName string) {
	words := cleanInput(str)

	commandName := words[0]
	command, exists := getNerfedCommands()[commandName]
	if exists {
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
		return commandName
	} else {
		return "Unknown command\n"
	}
}

func TestCmdExit(t *testing.T) {
	cases := []struct {
		input    string
		expected string
		testNum  int
	}{
		{
			input:    "exit",
			expected: "exit",
			testNum:  1,
		},
		{
			input:    "EXIT",
			expected: "exit",
			testNum:  2,
		},
		{
			input:    "quit",
			expected: "Unknown command\n",
			testNum:  3,
		},
		{
			input:    "ls",
			expected: "Unknown command\n",
			testNum:  4,
		},
	}

	for _, c := range cases {
		actual := startTestRepl(cfg, c.input)
		if actual != c.expected {
			t.Errorf("Test %d failed. \n Have: %v \n Want: %v", c.testNum, actual, c.expected)
		}
		fmt.Printf("Test %d passed. Command '%v' was successfully handled.", c.testNum, c.input)
	}
}

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
		testNum  int
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
			testNum:  1,
		},
		{
			input:    "HeLLo wORlD  ",
			expected: []string{"hello", "world"},
			testNum:  2,
		},
		{
			input:    " HeLLo woRlD",
			expected: []string{"hello", "world"},
			testNum:  3,
		},
		{
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
			testNum:  4,
		},
		{
			input:    "	HELLO WORLD	",
			expected: []string{"hello", "world"},
			testNum:  5,
		},
		{
			input:    "hello World	",
			expected: []string{"hello", "world"},
			testNum:  6,
		},
		{
			input:    "	hello World",
			expected: []string{"hello", "world"},
			testNum:  7,
		},
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
			testNum:  8,
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Test %d failed. \n Have: %s \n Want: %s", c.testNum, word, expectedWord)
			} else {
				fmt.Printf("Test %d passed. String '%s' was successfully cleaned", c.testNum, c.input)
			}
		}
	}
}
