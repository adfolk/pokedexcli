package main

import (
	"fmt"
	"os"
)

func cmdExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
