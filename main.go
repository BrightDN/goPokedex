package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedéx > ")

		if !scanner.Scan() {
			fmt.Println("Something went wrong, exiting program.")
			break
		}

		cleanInput := CleanInput(scanner.Text())
		if len(cleanInput) == 0 {
			fmt.Println("Empty command detected!")
			continue
		}

		fmt.Printf("Your command was: %s\n", cleanInput[0])
	}
}

func CleanInput(text string) []string {
	lowerCased := strings.ToLower(text)
	trimmed := strings.TrimSpace(lowerCased)
	return strings.Fields(trimmed)
}
