package pokeweb

import (
	"fmt"
	"io"
	"net/http"
)

func GetResource(url string) (response []byte, err error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Network error")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read in http response")
	}
	res.Body.Close()
	if res.StatusCode >= 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	return body, nil
}
