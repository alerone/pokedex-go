package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"
)

type AreaLocationsResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocations(pageURL *string) (AreaLocationsResponse, error){
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if res, ok := c.cache.Get(url); ok {
		fmt.Println("\tCached response")
		fmt.Println()
		locationsResponse := AreaLocationsResponse{}
		err := json.Unmarshal(res, &locationsResponse)
		if err != nil {
			return AreaLocationsResponse{}, err
		}
		return locationsResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaLocationsResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AreaLocationsResponse{}, err
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaLocationsResponse{}, err
	}

	locationsResponse := AreaLocationsResponse{}
	err = json.Unmarshal(dat, &locationsResponse)
	if err != nil {
		return AreaLocationsResponse{}, err
	}

	c.cache.Add(url, dat)

	return locationsResponse, nil
}
