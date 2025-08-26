package pokeapi

import (
	"fmt"
    "encoding/json"
	"net/http"
	"io"
)

func (c *Client) FetchEncounters(url string) (*PokemonEncountersResponse, error) {
    encounters := PokemonEncountersResponse{}
    if err := c.CheckCache(url, &encounters); err == nil {
        return &encounters, nil
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("Error creating request: %v", err)
    }
    
    res, err := c.httpClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("Error making request: %v", err)
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, fmt.Errorf("Error reading response: %v", err)
    }
    
    var response PokemonEncountersResponse
    if err := json.Unmarshal(body, &response); err != nil {
        return nil, fmt.Errorf("Error unmarshalling JSON: %v", err)
    }
    
    c.cache.Add(url, body)
    return &response, nil
}