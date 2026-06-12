package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(area string) (RespAreaInfo, error) {
	url := baseURL + "/location-area/" + area
	var data []byte

	val, exists := c.cache.Get(url)
	if exists {
		data = val
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespAreaInfo{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespAreaInfo{}, err
		}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return RespAreaInfo{}, err
	}
	c.cache.Add(url, data)
	}

	pokemonResp := RespAreaInfo{}
	err := json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return RespAreaInfo{}, err
	}

	return pokemonResp, nil
}