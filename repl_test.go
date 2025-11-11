package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "henlo  mfer",
			expected: []string{"henlo", "mfer"},
		},
		{
			input:    "	wtf mate",
			expected: []string{"wtf", "mate"},
		},
		{
			input:    "mentlegen ",
			expected: []string{"mentlegen"},
		},
		{
			input:    "hElLo wORlD ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Expected %v, got %v instead", expectedWord, word)
			}
		}
	}
}
