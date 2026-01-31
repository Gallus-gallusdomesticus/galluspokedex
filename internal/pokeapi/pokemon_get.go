package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon string) (PokemonType, error) {
	url := baseURL + "/pokemon/" + pokemon

	if data, ok := c.cache.Get(url); ok {
		var pokResp PokemonType
		if err := json.Unmarshal(data, &pokResp); err != nil {
			return PokemonType{}, err
		}
		return pokResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonType{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonType{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonType{}, err
	}

	var pokResp PokemonType
	if err := json.Unmarshal(data, &pokResp); err != nil {
		return PokemonType{}, err
	}
	c.cache.Add(url, data)

	return pokResp, nil

}
