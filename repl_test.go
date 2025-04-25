package main

import (
	"fmt"
	"testing"
)

var nerfedCommands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit, but does not actually call os.Exit(0)",
		callback:    nerfedCommandExit,
	},
}

func nerfedCommandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	return nil
}

func TestCmdExit(t *testing.T) {
	cases := []struct {
		input    string
		expected string
		isErr    bool
		testNum  int
	}{
		{
			input:    "exit",
			expected: "",
			isErr:    false,
			testNum:  1,
		},
		{
			input:    "EXIT",
			expected: "Unknown command\n",
			isErr:    true,
			testNum:  2,
		},
		{
			input:    "quit",
			expected: "Unknown command\n",
			isErr:    true,
			testNum:  3,
		},
		{
			input:    "ls",
			expected: "Unknown command\n",
			isErr:    true,
			testNum:  4,
		},
	}

	for _, c := range cases {
		actual := runCommand(c.input, nerfedCommands)
		if actual != nil && c.isErr == true {
			res := fmt.Sprintf("%v", actual)
			if res != c.expected {
				t.Errorf("Test %d failed. \n Have: %v \n Want: %v", c.testNum, res, c.expected)
			}
		} else {
			fmt.Printf("Test %d passed. Command '%v' was successfully handled.", c.testNum, c.input)
		}
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
