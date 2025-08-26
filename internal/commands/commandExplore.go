package commands

import (
	"fmt"
	"github.com/brightDN/goPokedex/internal/pokeapi"
)
func commandExplore(cfg *pokeapi.Config) error{
	url := cfg.BaseURL + "location-area/" + SupportedCommands["explore"].Args[0].Val
	
	response, err := cfg.PokeapiClient.FetchEncounters(url)
    if err != nil {
		fmt.Println("Invalid location")
        return err
    }

	fmt.Printf("Exploring %s...\n", response.Name)
	encounters := response.PokemonEncounters

	if len(encounters) == 0 {
		fmt.Println("No pokemons in this area...")
	} else {
		fmt.Println("Pokemon found:")
	}
	for _, encounter := range encounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}