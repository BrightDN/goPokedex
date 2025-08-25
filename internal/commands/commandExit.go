package commands

import (
	"fmt"
	"os"
	"github.com/brightDN/goPokedex/internal/pokeapi"
)

func commandExit(cfg *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}