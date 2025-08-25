package commands

import (
	"fmt"
	"os"

    "github.com/brightDN/goPokedex/internal/pokeapi"
)

var SupportedCommands map[string]CliCommand

func InitCommands() {
    SupportedCommands = map[string]CliCommand{
        "exit": {
            Name:        "exit",
            Description: "Exit the program",
            Callback:    commandExit,
            Args:        []CommandArgs{},
        },
        "help": {
            Name:        "help", 
            Description: "Displays a help message",
            Callback:    commandHelp,
            Args:        []CommandArgs{},
        },
        "map": {
            Name:        "map",
            Description: "Show the next 20 locations", 
            Callback:    commandMap,
            Args:        []CommandArgs{},
        },
		"mapb": {
			Name:           "mapb",
			Description:    "Show the previouw 20 locations",
			Callback:       commandMapB,
            Args:           []CommandArgs{},
		},
        "explore": {
            Name: "explore",
            Description: "Explore a specified area",
            Callback:     commandExplore,
            Args:        []CommandArgs{
                                {Name: "AreaName"},
                            },
        },
    }
}

func commandExit(cfg *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *pokeapi.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
    fmt.Println("")

	for _, command := range SupportedCommands {
        arguments := "This command expects no arguments\n"
        if len(command.Args) > 0 {
            arguments = "Expected args: "
            for _, arg := range command.Args {
                if len(command.Args) == 1 {
                    arguments += arg.Name
                    } else {
                    arguments += arg.Name + ", "
                }
            }
            arguments += "\n"
        }
		fmt.Printf("%s: %s\n%s\n", command.Name, command.Description, arguments)
	}
	return nil
}

func commandMap(cfg *pokeapi.Config) error {
    url := "https://pokeapi.co/api/v2/location-area/"
    if cfg.NextLocationURL != nil {
        url = *cfg.NextLocationURL
    }

    if cfg.NextLocationURL == nil && cfg.PreviousLocationURL != nil {
        fmt.Println("You are on the last page")
        return nil
    }

    response, err := cfg.PokeapiClient.FetchLocationAreas(url)
    if err != nil {
        return err
    }

    for _, location := range response.Results {
        fmt.Println(location.Name)
    }
    
    cfg.NextLocationURL = response.Next
    cfg.PreviousLocationURL = response.Previous
    return nil
}

func commandMapB(cfg *pokeapi.Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
    if cfg.PreviousLocationURL != nil {
        url = *cfg.PreviousLocationURL
    } else {
		fmt.Println("You are on the first page")
		return nil
	}

    response, err := cfg.PokeapiClient.FetchLocationAreas(url)
    if err != nil {
        return err
    }

    for _, location := range response.Results {
        fmt.Println(location.Name)
    }
    
    cfg.NextLocationURL = response.Next
    cfg.PreviousLocationURL = response.Previous
    return nil
}