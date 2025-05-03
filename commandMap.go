package main

import (
	"errors"
	"fmt"
	"github.com/adfolk/bootdev-guided/pokedexcli/internal/pokeapi"
)

func commandMapf(cfg *config) error {
	if cfg.nextLocationsURL == nil {
		newUrl, rawLocations, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL)
		if err != nil {
			return err
		}
		cfg.pokecacheCache.Add(newUrl, rawLocations)
		respLocs, err := pokeapi.UnmarshallLocations(rawLocations)

		cfg.nextLocationsURL = respLocs.Next
		cfg.prevLocationsURL = respLocs.Previous

		for _, loc := range respLocs.Results {
			fmt.Println(loc.Name)
		}
		return nil
	}
	rawCachedLocs, exists := cfg.pokecacheCache.Get(*cfg.nextLocationsURL)
	if !exists {
		newUrl, rawLocations, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL)
		if err != nil {
			return err
		}
		cfg.pokecacheCache.Add(newUrl, rawLocations)
		respLocs, err := pokeapi.UnmarshallLocations(rawLocations)

		cfg.nextLocationsURL = respLocs.Next
		cfg.prevLocationsURL = respLocs.Previous

		for _, loc := range respLocs.Results {
			fmt.Println(loc.Name)
		}
		return nil

	}

	parsedLocs, err := pokeapi.UnmarshallLocations(rawCachedLocs)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = parsedLocs.Next
	cfg.prevLocationsURL = parsedLocs.Previous
	for _, loc := range parsedLocs.Results {
		fmt.Println(loc.Name)
	}
	//fmt.Println("\nIT WORKED!!!!\n")
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	rawCachedLocs, exists := cfg.pokecacheCache.Get(*cfg.prevLocationsURL)
	if !exists {
		newUrl, rawLocations, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationsURL)
		if err != nil {
			return err
		}
		cfg.pokecacheCache.Add(newUrl, rawLocations)
		respLocs, err := pokeapi.UnmarshallLocations(rawLocations)

		cfg.nextLocationsURL = respLocs.Next
		cfg.prevLocationsURL = respLocs.Previous

		for _, loc := range respLocs.Results {
			fmt.Println(loc.Name)
		}
		return nil
	}

	parsedLocs, err := pokeapi.UnmarshallLocations(rawCachedLocs)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = parsedLocs.Next
	cfg.prevLocationsURL = parsedLocs.Previous

	for _, loc := range parsedLocs.Results {
		fmt.Println(loc.Name)
	}
	//fmt.Println("\nIT WORKED!!!!\n")
	return nil
}
