package commands

import(
	"fmt"
	"github.com/brightDN/goPokedex/internal/pokeapi"
)

func commandMapB(cfg *pokeapi.Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
    if cfg.PreviousLocationURL != nil {
        url = *cfg.PreviousLocationURL
    } else {
		fmt.Println("You are on the first page")
		return nil
	}

    response, err := cfg.PokeapiClient.FetchLocationAreas(url)
    if err != nil {
        return err
    }

    for _, location := range response.Results {
        fmt.Println(location.Name)
    }
    
    cfg.NextLocationURL = response.Next
    cfg.PreviousLocationURL = response.Previous
    return nil
}