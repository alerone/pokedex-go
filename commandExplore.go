package main

import "fmt"

func commandExplore(cfg *config, params []string) error {
	location := params[0]
	if location == "" {
		return fmt.Errorf("missing location in request: explore <location>")
	}

	pokemons, err := cfg.pokeapiClient.Explore(location)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")
	for _,pokemon := range pokemons {
		fmt.Printf("- %s\n", pokemon)
	}

	return nil
}



