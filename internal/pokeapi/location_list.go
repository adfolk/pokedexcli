package pokeapi

import (
	"encoding/json"
)

func UnmarshallLocations(dat []byte) (RespShallowLocations, error) {
	locationsResp := RespShallowLocations{}
	err := json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}

func (c *Client) GetLocations(pageURL *string) (string, []byte, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	dat, err := c.GetResource(url)
	if err != nil {
		return "", nil, err
	}
	return url, dat, nil
}
