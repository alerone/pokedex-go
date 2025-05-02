package main

import (
	"fmt"
)

func commandInspect(cfg *config, params []string) error {
	pokemonReq := ""
	force := ""

	if len(params) > 0 {
		pokemonReq = params[0]
		if len(params) > 1 {
			force = params[1]
		}
	}


	if pokemonReq == "" {
		return fmt.Errorf("specify a pokemon to inspect. inspect <pokemon>")
	}

	pokemon, ok := cfg.pokedex[pokemonReq]

	if !ok {
		fmt.Println(force)
		return fmt.Errorf("you have not caught this pokemon...")
	}

	fmt.Println()
	fmt.Printf("Pokedex ID: %d\n", pokemon.ID)
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	return nil
}
