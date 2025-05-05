package main

import (
	"bytes"
	"fmt"

	"github.com/alerone/pokedex-go/internal/pokeapi"
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
	fmt.Printf("Height: %.2f m\n", float32(pokemon.Height)/10.0)  // decimeters to meters
	fmt.Printf("Weight: %.2f kg\n", float32(pokemon.Weight)/10.0) // hectograms to kg
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	fmt.Println()
	fmt.Println("Evolution Chain")
	fmt.Println("---------------")
	evoChain := displayEvolutionChain(pokemon.EvoChain)
	fmt.Println(evoChain)

	return nil
}

func displayEvolutionChain(evoChain *pokeapi.EvolutionChainLink) string {
	if evoChain == nil {
		return "\n"
	}
	var out bytes.Buffer

	if len(evoChain.EvolutionDetails) == 0 {
		fmt.Fprint(&out, "  - ")
	}
	if len(evoChain.EvolutionDetails) > 0 {
		fmt.Fprintf(&out, "-> %s= ", evoChain.EvolutionDetails[0].String())
	}
	fmt.Fprintf(&out, "%s ", evoChain.Species.Name)

	for _, evo := range evoChain.EvolvesTo {
		outs := out.String()
		if len(outs) > 0 {
			if outs[len(outs)-1] == '\n' {
				fmt.Fprint(&out, "\t")
			}
		}
		if len(evoChain.EvolvesTo) > 1 {
			fmt.Fprintf(&out, "- %s", displayEvolutionChain(evo))
		} else {
			fmt.Fprintf(&out, "%s", displayEvolutionChain(evo))
		}
	}
	fmt.Fprintln(&out)

	return out.String()
}
