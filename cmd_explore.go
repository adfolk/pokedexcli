package main

import "fmt"

func cmdExplore(cfg *config, locName ...string) error {
	res, err := cfg.pokeapiClient.ListPokemon(locName[0])
	if err != nil {
		return err
	}
	for _, val := range res.PokemonEncounters {
		fmt.Println(val.Pokemon.Name)
	}
	return nil
}
