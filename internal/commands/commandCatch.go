package commands

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/brightDN/goPokedex/internal/pokeapi"
)

func commandCatch(cfg *pokeapi.Config) error {

	url := cfg.BaseURL + "pokemon/" + SupportedCommands["catch"].Args[0].Val
	resp, err := cfg.PokeapiClient.FetchPokemon(url)

	if err != nil {
		fmt.Println("Wrong parameter")
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", resp.Name)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	if isCaught := tryCatch(resp.BaseExperience, random); isCaught {
		fmt.Printf("%s was caught!\n", resp.Name)
	
		entry := pokeapi.PokedexEntry{
		IsCaught:   true,
		Data:   *resp,
		}
	
		cfg.UserPokedex[resp.Name] = entry
	} else {
		fmt.Printf("%s Escaped...\n", resp.Name)
	}

	return nil
}



func catchChance(baseExp int) float64 {
	const minExp = 50
	const maxExp = 400

	if baseExp < minExp {
		baseExp = minExp
	}
	if baseExp > maxExp {
		baseExp = maxExp
	}

	return 0.9 - 0.8*float64(baseExp-minExp)/float64(maxExp-minExp)
}

func tryCatch(baseExp int, r *rand.Rand) bool {
	return r.Float64() < catchChance(baseExp)
}