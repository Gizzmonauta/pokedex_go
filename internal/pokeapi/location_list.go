package pokeapi

import (
	"encoding/json"
	"net/http"
)

// ListLocations - 
func (c *Client) ListLocations(pageURL *string) (RespLocationAreas, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationAreas{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationAreas{}, err
	}
	defer res.Body.Close()

	respLocationAreas := RespLocationAreas{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&respLocationAreas)
	if err != nil {
		return RespLocationAreas{}, err
	}

	return respLocationAreas, nil
}