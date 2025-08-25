package commands

import (
	"github.com/brightDN/goPokedex/internal/pokeapi"
)

type CliCommand struct {
	Name			string
	Description 	string
	Callback 		func(*pokeapi.Config) error
	Args 			[]CommandArgs
}

type CommandArgs struct {
	Name string
    Val string
}

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
                                {Name: "AreaName",
                                Val: "TempTest",},
                            },
        },
    }
}