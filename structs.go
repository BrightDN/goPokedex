package main

type PokeLocation struct {
	Name	string `json:"name"`
	Url		string `json:"url"`
}

type LocationAreaResponse struct {
    Results 	[]PokeLocation	`json:"results"`
	Next 		*string 		`json:"next"`
	Previous	*string			`json:"previous"`
}

type cliCommand struct {
	name			string
	description 	string
	callback 		func(*config) error
}

type config struct {
    nextLocationURL     *string
    previousLocationURL *string
}