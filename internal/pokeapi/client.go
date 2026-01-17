package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient *http.Client
}

//New Client for PokeApi

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &Client{
		httpClient: httpClient,
	}
}

// ListLocations fetches one page of location-areas
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	var locationsResp RespShallowLocations
	if err := json.Unmarshal(data, &locationsResp); err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
