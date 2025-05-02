package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Explore(location string) ([]string, error) {
	result := make([]string, 0)
	var err error

	if location == "" {
		return result, fmt.Errorf("location not provided")
	}
	url := fmt.Sprintf("%s/location-area/%s", baseURL, location)

	data, ok := c.cache.Get(url)

	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return result, err
		}
		res, err := c.httpClient.Do(req)
		if err != nil {
			return result, err
		}
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return result, err
		}
		c.cache.Add(url, data)
	} else {
		fmt.Printf("\tCached result\n")
	}

	var fullResponse encountersResponse
	err = json.Unmarshal(data, &fullResponse)
	if err != nil {
		return result, err
	}

	for _, pokemon := range fullResponse.PokemonEncounters {
		result = append(result, pokemon.Pokemon.Name)
	}

	return result, nil
}
