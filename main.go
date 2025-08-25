package main

import (
	"fmt"
	"strings"
	"time"
	"bufio"
	"os"

	"github.com/brightDN/goPokedex/internal/pokeapi"
	"github.com/brightDN/goPokedex/internal/commands"
)

func main() {
	commands.InitCommands()
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &pokeapi.Config{
		PokeapiClient: pokeClient,
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
				command.Callback(cfg)
				break
			}
		}
		if !commandFound {
			fmt.Println("Unknown command")
		}
	}
}

func CleanInput(text string) []string {
	if len(text) == 0 {
        return []string{""}
    }
	lowerCased := strings.ToLower(text)
	trimmed := strings.TrimSpace(lowerCased)
	return strings.Fields(trimmed)
}