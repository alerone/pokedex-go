package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokemonSpeciesResult struct {
	EvolutionChainURL struct {
		URL string `json:"url"`
	} `json:"evolution_chain"`
}

func (c *Client) GetEvolutionChain(pokemonID string) (EvolutionChainLink, error) {
	url, err := c.getEvolutionChainFromSpecies(pokemonID)
	var evoChain evoChainRes

	if err != nil {
		return EvolutionChainLink{}, err	
	}

	if data, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(data, &evoChain); err != nil {
			return EvolutionChainLink{}, err	
		}

		return *evoChain.Chain, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return EvolutionChainLink{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return EvolutionChainLink{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return EvolutionChainLink{}, err
	}

	c.cache.Add(url, data)

	if err := json.Unmarshal(data, &evoChain); err != nil {
		return EvolutionChainLink{}, err
	}

	return *evoChain.Chain, nil
}

func (c *Client) getEvolutionChainFromSpecies(pokemonID string) (string, error) {
	url := fmt.Sprintf("%s/pokemon-species/%s", baseURL, pokemonID)
	var species pokemonSpeciesResult

	if data, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(data, &species); err != nil {
			return "", nil
		}

		return species.EvolutionChainURL.URL, nil
	}

	if pokemonID == "" {
		return "", fmt.Errorf("specify a pokemon")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	c.cache.Add(url, data)

	if err = json.Unmarshal(data, &species); err != nil {
		return "", err
	}

	return species.EvolutionChainURL.URL, nil
}
