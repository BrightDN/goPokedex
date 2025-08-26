package main

import (
	"fmt"
	"time"
	"bufio"
	"os"

	"github.com/brightDN/goPokedex/internal/pokeapi"
	"github.com/brightDN/goPokedex/internal/commands"
	"github.com/brightDN/goPokedex/internal/pokedex"
)

func main() {
	userPokedex := map[string]pokedex.PokedexEntry{}
	commands.InitCommands()
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &pokeapi.Config{
		PokeapiClient: pokeClient,
		BaseURL: "https://pokeapi.co/api/v2/",
		UserPokedex: userPokedex,
	} 

	scanner := bufio.NewScanner(os.Stdin)
	commandFound := false

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			fmt.Println("Something went wrong, exiting program.")
			break
		}

		cleanInput := CleanInput(scanner.Text())

		commandFound = false

		for _, command := range commands.SupportedCommands {
			if cleanInput[0] == command.Name {
				commandFound = true
				cmd := commands.SupportedCommands[cleanInput[0]]
				if len(cleanInput) < (len(cmd.Args) + 1) {
					fmt.Println("Missing arguments")
					} else {
						for id := range cmd.Args {
							cmd.Args[id].Val = cleanInput[id+1]
						}
					commands.SupportedCommands[cleanInput[0]] = cmd
					command.Callback(cfg)
					break
				}
				}
			}
		if !commandFound {
			fmt.Println("Unknown command")
		}
	}
}