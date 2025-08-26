package commands

import(
	"fmt"
	"github.com/brightDN/goPokedex/internal/pokeapi"
)

func commandMap(cfg *pokeapi.Config) error {
    url := cfg.BaseURL + "/location-area/"
    if cfg.NextLocationURL != nil {
        url = *cfg.NextLocationURL
    }

    if cfg.NextLocationURL == nil && cfg.PreviousLocationURL != nil {
        fmt.Println("You are on the last page")
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