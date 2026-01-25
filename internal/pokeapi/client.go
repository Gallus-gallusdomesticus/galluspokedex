package pokeapi

import (
	"net/http"
)

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
