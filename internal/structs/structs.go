package structs

import (
	"net/http"
	"github.com/brightDN/goPokedex/internal/cache"
	"github.com/brightDN/goPokedex/internal/pokeapi"
)

type CommandArgs struct {
	Name string
}

type Config struct {
	NextLocationURL     *string
    PreviousLocationURL *string
	PokeapiClient    	pokeapi.Client
}

type CliCommand struct {
	Name			string
	Description 	string
	Callback 		func(*Config) error
	Args 			[]CommandArgs
}

type Client struct {
	Cache      cache.Cache
	HttpClient http.Client
}