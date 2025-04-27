package main

import (
	"encoding/json"
	"fmt"
	"github.com/adfolk/bootdev-guided/pokedexcli/internal/pokeweb"
)

type mapPage struct {
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
}

var mapConfig = config{
	Previous: "",
	Next:     "https://pokeapi.co/api/v2/location-area/",
}

func getLocations(c *config) error {
	url := c.Next
	locs, err := pokeweb.GetResource(url)
	if err != nil {
		return err
	}

	var stashedLocs mapPage
	err = json.Unmarshal(locs, &stashedLocs)
	if err != nil {
		return err
	}

	for _, loc := range stashedLocs.Results {
		fmt.Println(loc.Name)
	}
	err = updateConfig(c, locs)
	if err != nil {
		return err
	}
	return nil
}

func getPrevLocations(c *config) error {
	url := c.Previous
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	locs, err := pokeweb.GetResource(url)
	if err != nil {
		return err
	}

	var stashedLocs mapPage
	err = json.Unmarshal(locs, &stashedLocs)
	if err != nil {
		return err
	}

	for _, loc := range stashedLocs.Results {
		fmt.Println(loc.Name)
	}
	err = updateConfig(c, locs)
	if err != nil {
		return err
	}
	return nil
}
