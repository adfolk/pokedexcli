package main

import (
	"time"

	"github.com/adfolk/bootdev-guided/pokedexcli/internal/pokeapi"
	"github.com/adfolk/bootdev-guided/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCacheDuration := 30 * time.Second
	locsCache := pokecache.NewCache(pokeCacheDuration)

	cfg := &config{
		pokeapiClient:  pokeClient,
		pokecacheCache: locsCache,
	}
	startRepl(cfg)
}
