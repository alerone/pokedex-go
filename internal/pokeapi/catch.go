package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const maxIntents = 15

func (c *Client) Catch(pokeURI string) (Pokemon, bool, error) {
	url := fmt.Sprintf("%s/pokemon/%s", baseURL, strings.ToLower(pokeURI))
	var pokemon Pokemon

	if data, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return Pokemon{}, false, err
		}
		catchRes := tryCatch(pokemon, maxIntents)
		return pokemon, catchRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, false, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, false, fmt.Errorf("%s not found", pokeURI)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, false, fmt.Errorf("%s not found", pokeURI)
	}

	c.cache.Add(url, data)

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, false, fmt.Errorf("%s not found", pokeURI)
	}

	catchRes := tryCatch(pokemon, maxIntents)


	return pokemon, catchRes, nil
}
