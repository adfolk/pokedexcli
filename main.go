package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	results := strings.Fields(strings.ToLower(text))
	return results
}
