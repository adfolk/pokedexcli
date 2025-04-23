package main

import (
	"fmt"
	"testing"
)

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
