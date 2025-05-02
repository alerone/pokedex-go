package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alerone/pokedex-go/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, params []string) error
}

type config struct {
	pokeapiClient       pokeapi.Client
	PreviousLocationURL *string
	NextLocationURL     *string
	pokedex             map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	commands := getCommands()
	var userInput string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		userInput = scanner.Text()
		tokens := cleanInput(userInput)
		if cmd, ok := commands[tokens[0]]; ok {
			err := cmd.callback(cfg, tokens[1:])
			if err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
		fmt.Print("Pokedex > ")
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location and know which pokemon could appear in that location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon and save it to the Pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon entry in the pokedex",
			callback:    commandInspect,
		},
	}
}
