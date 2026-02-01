package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetLocation fetches one page of pokemon in location-areas

func (c *Client) GetLocation(area string) (LocationType, error) {

	url := baseURL + "/location-area/" + area

	if data, ok := c.cache.Get(url); ok { //cache checking if its exist already
		var locResp LocationType
		if err := json.Unmarshal(data, &locResp); err != nil {
			return LocationType{}, err
		}
		return locResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationType{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationType{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationType{}, err
	}

	var pokeResp LocationType
	if err := json.Unmarshal(data, &pokeResp); err != nil {
		return LocationType{}, err
	}

	c.cache.Add(url, data) //add cache on http request

	return pokeResp, nil
}
