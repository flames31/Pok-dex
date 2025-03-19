package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	pokemonResp := RespPokemon{}

	if resp, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(resp, &pokemonResp)
		if err != nil {
			return RespPokemon{}, err
		}
		fmt.Println("RETRIEVED FROM CACHE")
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return RespPokemon{}, fmt.Errorf("cannot find pokemon")
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, data)

	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return RespPokemon{}, err
	}

	return pokemonResp, nil
}
