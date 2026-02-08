package pokeapi

import (
	"net/http"
	"time"
)

// Client - 
type Client struct {
	httpClient http.Client
}

// NewClient - This is a "Constructor" function. 
// It creates a client with a 1-minute timeout.
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
