package main

import (
	"fmt"
	"os"
)
	

var supportedCommands map[string]cliCommand

func initCommands() {
    supportedCommands = map[string]cliCommand{
        "exit": {
            name:        "exit",
            description: "Exit the program",
            callback:    commandExit,
        },
        "help": {
            name:        "help", 
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "map": {
            name:        "map",
            description: "Show the next 20 locations", 
            callback:    commandMap,
        },
		"mapb": {
			name: "mapb",
			description: "Show the previouw 20 locations",
			callback: commandMapB,
		},
    }
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
    fmt.Println("")

	for _, command := range supportedCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config) error {
    url := "https://pokeapi.co/api/v2/location-area/"
    if cfg.nextLocationURL != nil {
        url = *cfg.nextLocationURL
    }

    if cfg.nextLocationURL == nil && cfg.previousLocationURL != nil {
        fmt.Println("You are on the last page")
        return nil
    }

    response, err := cfg.pokeapiClient.fetchLocationAreas(url)
    if err != nil {
        return err
    }

    for _, location := range response.Results {
        fmt.Println(location.Name)
    }
    
    cfg.nextLocationURL = response.Next
    cfg.previousLocationURL = response.Previous
    return nil
}

func commandMapB(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
    if cfg.previousLocationURL != nil {
        url = *cfg.previousLocationURL
    } else {
		fmt.Println("You are on the first page")
		return nil
	}

    response, err := cfg.pokeapiClient.fetchLocationAreas(url)
    if err != nil {
        return err
    }

    for _, location := range response.Results {
        fmt.Println(location.Name)
    }
    
    cfg.nextLocationURL = response.Next
    cfg.previousLocationURL = response.Previous
    return nil
}
