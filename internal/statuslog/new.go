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

func NewStatus(status string) (bool, error){
	if status == "" {
		return false, errors.New("no status provided")
	}

	address, err := env.GetEnvValue("ADDRESS", "")

	// Check for errors getting env var
	if err != nil {
		return false, err
	}

	// Format target for POST request
	target := fmt.Sprintf("https://api.omg.lol/address/%s/statuses/", address)

	// Format data for POST request
	// data := fmt.Sprintf(`{"status": "%s"}`, status)
	data := map[string]string{"status": status}
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

	// Create new request
	req, err := http.NewRequest("POST", target, bytes.NewBuffer(jsonData))
	if err != nil {
		return false, err
	}

	// Add auth headers to request
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	// Close response body when finished
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result NewStatusData
	if err := json.Unmarshal(body, &result); err != nil {
		return false, errors.New("could not unmarshal JSON")
	}

	if result.Request.StatusCode != 200 {
		errorMsg := fmt.Sprintf("status code: %d", result.Request.StatusCode)
		return false, errors.New(errorMsg)
	}

	return true, nil
}
