package main

import (
	"net/http"
	"time"
	"io"
	"fmt"
	"encoding/json"
)

func commandMap(cfg *config) error{

	url := "https://pokeapi.co/api/v2/location-area/"
    if cfg.nextLocationURL != nil {
        url = *cfg.nextLocationURL
    }

	if cfg.nextLocationURL == nil && cfg.previousLocationURL != nil {
		fmt.Println("You are on the last page")
		return nil
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
    	return fmt.Errorf("Error getting the requested url: %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
    	return fmt.Errorf("Error handling the request: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
    	return fmt.Errorf("Error reading JSON response: %v", err)
	}
	
	var response LocationAreaResponse
	if err := json.Unmarshal(body, &response); err != nil {
    	return fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	for _, location := range response.Results {
    	fmt.Println(location.Name)
	}

    cfg.nextLocationURL = response.Next
    cfg.previousLocationURL = response.Previous
	return nil
}

type PokeLocation struct
{
	Name string `json:"name"`
	Url string `json:"url"`
}

type LocationAreaResponse struct {
    Results 	[]PokeLocation	`json:"results"`
	Next 		*string 		`json:"next"`
	Previous	*string			`json:"previous"`
}