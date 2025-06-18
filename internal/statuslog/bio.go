package statuslog

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"internal/env"
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

func BioEdit(bio string) (bool, error) {
	// Confirm bio is provided
	if bio == "" {
		return false, errors.New("no bio provided")
	}

	// Get address
	address, err := env.GetEnvValue("ADDRESS", "")
	if err != nil {
		return false, err
	}

	// Format target for POST request
	target := fmt.Sprintf("https://api.omg.lol/address/%s/statuses/bio/", address)

	// Format data for POST request
	data := map[string]string{"content": bio}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	// Get api key for bearer token
	apiKey, err := env.GetEnvValue("API_KEY", "")
	if err != nil {
		return false, err
	}

	// Format auth bearer token (header)
	bearer := "Bearer " + apiKey

	// Create a new request
	req, err := http.NewRequest("POST", target, bytes.NewBuffer(jsonData))
	if err != nil {
		return false, err
	}

	// Add auth header to request
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	// Close resposne body when finished
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// Convert it into NewBioData type
	var result NewBioData
	if err := json.Unmarshal(body, &result); err != nil {
		return false, errors.New("could not unmarshal JSON")
	}

	// Confirm correct status code
	if result.Request.StatusCode != 200 {
		errorMsg := fmt.Sprintf("status code: %d", result.Request.StatusCode)
		return false, errors.New(errorMsg)
	}

	return true, nil
}
