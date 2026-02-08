package pokeapi

// RespLocationAreas is the shape of the data returned 
// from the /location-area endpoint.
type RespLocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`     // We use *string because these can be null
	Previous *string `json:"previous"` // and pointers handle nulls perfectly in Go
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}