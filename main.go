package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	initCommands()
	cfg := &config{} 
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
		for _, command := range supportedCommands {
			if cleanInput[0] == command.name {
				commandFound = true
				command.callback(cfg)
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

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for _, command := range supportedCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

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
    }
}