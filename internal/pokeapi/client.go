package pokeapi

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
    "github.com/brightDN/goPokedex/internal/cache"
)

type Client struct {
    cache      cache.Cache
    httpClient http.Client
}

func (c *Client) FetchLocationAreas(url string) (*LocationAreaResponse, error) {
    
    locations := LocationAreaResponse{}
    if err := c.CheckCache(url, &locations); err == nil {
        return &locations, nil
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
    
    var response LocationAreaResponse
    if err := json.Unmarshal(body, &response); err != nil {
        return nil, fmt.Errorf("Error unmarshalling JSON: %v", err)
    }
    
    c.cache.Add(url, body)
    return &response, nil
}

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

func (c *Client) CheckCache(key string, target interface{}) error {
    val, ok := c.cache.Get(key)
    if !ok {
        return fmt.Errorf("cache miss for key: %s", key)
    }

    err := json.Unmarshal(val, target)
    if err != nil {
        return fmt.Errorf("error unmarshalling cached data: %w", err)
    }

    return nil 
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: cache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}