package main

import (
	"net/http"
	"github.com/brightDN/goPokedex/internal"
)

type PokeLocation struct {
	Name	string `json:"name"`
	Url		string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type cliCommand struct {
	name			string
	description 	string
	callback 		func(*config) error
}

type config struct {
    nextLocationURL     *string
    previousLocationURL *string
	pokeapiClient    	Client
}

type Client struct {
	cache      internal.Cache
	httpClient http.Client
}