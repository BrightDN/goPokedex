package pokeapi

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
    "github.com/brightDN/goPokedex/internal/cache"
)

type Client struct {
    cache      cache.Cache
    httpClient http.Client
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