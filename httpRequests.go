package main

import (
	"net/http"
	"io"
	"fmt"
	"encoding/json"
)

func (c *Client) fetchLocationAreas(url string) (*LocationAreaResponse, error) {
    
    	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationAreaResponse{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return &locationsResp, err
		}

		return &locationsResp, nil
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

