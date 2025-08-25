package commands

import (
	"github.com/brightDN/goPokedex/internal/pokeapi"
)

type CliCommand struct {
	Name			string
	Description 	string
	Callback 		func(*pokeapi.Config) error
	Args 			[]CommandArgs
}

type CommandArgs struct {
	Name string
}