package pokeapi

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io"
)

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
