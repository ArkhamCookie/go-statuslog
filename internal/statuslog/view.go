package statuslog

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Get a status from a given address and status
func StatusGet(address, status string) (*StatuslogRetrieveData, error) {
	if address == "" {
		return nil, errors.New("no address given")
	}

	target := fmt.Sprintf("https://api.omg.lol/address/%s/statuses/%s", address, status)
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

	var result StatuslogRetrieveData
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("could not unmarshal JSON")
	}

	if result.Request.StatusCode != 200 {
		errorMsg := fmt.Sprintf("status code: %d", result.Request.StatusCode)
		return &result, errors.New(errorMsg)
	}

	return &result, nil
}
