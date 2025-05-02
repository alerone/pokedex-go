package main

import "fmt"

func commandMapf(cfg *config, params []string) error {
	response, err := cfg.pokeapiClient.GetLocations(cfg.NextLocationURL)
	if err != nil {
		return err
	}
	
	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	cfg.NextLocationURL = &response.Next
	cfg.PreviousLocationURL = &response.Previous

	return nil
}

func commandMapb(cfg *config, params []string) error {
	if cfg.PreviousLocationURL == nil || *cfg.PreviousLocationURL == ""{
		return fmt.Errorf("You are on the first page!")
	}

	response, err := cfg.pokeapiClient.GetLocations(cfg.PreviousLocationURL)
	if err != nil {
		return err
	}
	
	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	cfg.NextLocationURL = &response.Next
	cfg.PreviousLocationURL = &response.Previous

	return nil
}
