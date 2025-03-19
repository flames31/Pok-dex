package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(locationName string) (RespExploreLocation, error) {
	url := baseURL + "/location-area/" + locationName
	exploreLocResp := RespExploreLocation{}

	if resp, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(resp, &exploreLocResp)
		if err != nil {
			return RespExploreLocation{}, err
		}
		fmt.Println("RETRIEVED FROM CACHE")
		return exploreLocResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespExploreLocation{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespExploreLocation{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return RespExploreLocation{}, fmt.Errorf("Invalid location!")
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespExploreLocation{}, err
	}

	c.cache.Add(url, data)

	err = json.Unmarshal(data, &exploreLocResp)
	if err != nil {
		return RespExploreLocation{}, err
	}

	return exploreLocResp, nil
}
