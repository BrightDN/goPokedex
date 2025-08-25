package main

import (
	"fmt"
	"strings"
	"time"
	"bufio"
	"os"
	"net/http"
	
	"github.com/brightDN/goPokedex/internal"
)

func main() {
	initCommands()
	pokeClient := NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
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

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: internal.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}