package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		if err := s.Err(); err != nil {
			resErr := fmt.Errorf("Error scanning user input: %w", err)
			fmt.Printf("%v", resErr)
		}
		t := s.Text()
		cleaned := cleanInput(t)
		fmt.Printf("Your command was: %v\n", cleaned[0])
	}
}

func cleanInput(text string) []string {
	results := strings.Fields(strings.ToLower(text))
	return results
}
