package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func mapf(c *Config) error {
	locs, err := getLocs(c.Next)
	fmt.Printf("Getting resource from: %s\n", c.Next)
	for _, res := range locs.Results {
		fmt.Printf("%s\n", res.Name)
	}

	fmt.Printf("Next page: %s\n", locs.Next)
	fmt.Printf("Previous page: %s\n", locs.Prev)
	c.Previous = locs.Prev
	c.Next = locs.Next

	return err
}

func mapb(c *Config) error {
	locs, err := getLocs(c.Previous)
	fmt.Printf("Getting resource from: %s\n", c.Previous)
	for _, res := range locs.Results {
		fmt.Printf("%s\n", res.Name)
	}

	fmt.Printf("Next page: %s\n", locs.Next)
	fmt.Printf("Previous page: %s\n", locs.Prev)
	c.Previous = locs.Prev
	c.Next = locs.Next

	return err
}

type PokeLocs struct {
	Next    string `json:"next"`
	Prev    string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
}

func getLocs(endPoint string) (PokeLocs, error) {
	res, err := http.Get(endPoint)
	results := PokeLocs{}
	if err != nil {
		return results, fmt.Errorf("Problem calling api: %v", err)
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&results); err != nil {
		return results, err
	}

	return results, nil
}
