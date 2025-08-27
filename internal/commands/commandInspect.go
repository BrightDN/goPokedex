package commands

import (
	"fmt"
	"github.com/brightDN/goPokedex/internal/pokeapi"
)

func commandInspect(cfg *pokeapi.Config) error {
	if val, ok := cfg.UserPokedex[SupportedCommands["inspect"].Args[0].Val]; ok {
		data := val.Data
		fmt.Printf("Name: %s\n", data.Name)
		fmt.Printf("Height: %d\n", data.Height)
		fmt.Printf("Weight: %d\n", data.Weight)
		fmt.Println("Stats:")
		stats := data.Stats

		for _, stat := range stats {
			name := stat.Stat.Name
			val := stat.BaseStat
			fmt.Printf("    - %s: %d\n", name, val)
		}

		types := data.Types

		fmt.Println("Types:")
		for _, pType := range types {
			fmt.Printf("    - %s\n", pType.Type.Name)
		}

	} else {
		fmt.Println("You have not caught that pokemon")
	}
	return nil
}