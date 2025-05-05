package main

import (
	"fmt"
)

func commandCatch(cfg *config, params []string) error {
	pokeURI := ""
	if len(params) > 0 {
		pokeURI = params[0]
	}
	if pokeURI == "" {
		return fmt.Errorf("no pokemon provided. Try: catch <pokemon name or pokedex id>")
	}

	pokemon, resCatch, err := cfg.pokeapiClient.Catch(pokeURI)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if resCatch {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		evoChain, err := cfg.pokeapiClient.GetEvolutionChain(pokeURI)
		if err != nil {
			return err
		}

		pokemon.EvoChain = &evoChain

		if _, ok := cfg.pokedex[pokemon.Name]; !ok {
			cfg.pokedex[pokemon.Name] = pokemon
		}

	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

