package commands

import (
	"fmt"
	"github.com/brightDN/goPokedex/internal/pokeapi"
)

func commandPokedex (cfg *pokeapi.Config) error {
	if len(cfg.UserPokedex) == 0 {
		fmt.Println("You haven't caught any pokemon yet")
		return nil
	}
	
	fmt.Println("Your pokedex:")
	entries := cfg.UserPokedex

	for _, entry := range entries {
		fmt.Printf("    - %s\n", entry.Data.Name)
	}
	return nil
}