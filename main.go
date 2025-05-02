package main

import (
	"time"

	"github.com/alerone/pokedex-go/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(10 * time.Second, 10 * time.Second)
	cfg := &config{
		pokedex: make(map[string]pokeapi.Pokemon),
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
