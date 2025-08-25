package commands

import (
	"fmt"
	"github.com/brightDN/goPokedex/internal/pokeapi"
)
func commandExplore(cfg *pokeapi.Config) error{
	fmt.Println(SupportedCommands["explore"].Args[0].Val)
	return nil
}