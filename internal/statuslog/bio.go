package statuslog

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func BioGet(address string) (*StatuslogBioData, error) {
	if address == "" {
		return nil, errors.New("no address given")
	}

	target := fmt.Sprintf("https://api.omg.lol/address/%s/statuses/bio", address)
	resp, err := http.Get(target)
	if err != nil {
		errorMsg := fmt.Sprintf("could not create GET request: %s", err)
		return nil, errors.New(errorMsg)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result StatuslogBioData
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("could not unmarshal JSON")
	}

	if result.Request.StatusCode != 200 {
		errorMsg := fmt.Sprintf("status code: %d", result.Request.StatusCode)
		return &result, errors.New(errorMsg)
	}

	return &result, nil
}
