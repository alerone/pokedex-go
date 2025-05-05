package main

import "fmt"

func commandPokedex(cfg *config, params []string) error {
	captured := len(cfg.pokedex)
	if captured <= 0 {
		return fmt.Errorf("you have not captured any pokemon!")
	}
	fmt.Printf("you have captured %d pokemon!\n", captured)
	fmt.Println()

	for _, pokemon := range cfg.pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
