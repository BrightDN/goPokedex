package commands

import (
	"fmt"
	"github.com/brightDN/goPokedex/internal/pokeapi"
)
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