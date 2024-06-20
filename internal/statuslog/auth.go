package statuslog

import (
	"encoding/json"
	"errors"
	"fmt"
	"internal/env"
	"io"
	"net/http"
)

// Confirm auth status by getting address name
func AuthStatus(email string) (string, error) {
	// Check if an email is not given
	if email == "" {
		// If email is not given, get email from env file
		newemail, err := env.GetEnvValue("EMAIL", "")

		// Make sure email is gotten
		if newemail == "" {
			return "", errors.New("no email found")
		}
		if err != nil {
			return "", err
		}
		// Set email to env file's email
		email = newemail
	}

	// Get api key for bearer token
	apiKey, err := env.GetEnvValue("API_KEY", "")
	if err != nil {
		return "", err
	}

	// Format target for GET request
	target := fmt.Sprintf("https://api.omg.lol/account/%s/name", email)

	// Format auth bearer token (header)
	bearer := "Bearer " + apiKey

	// Create new request
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return "", err
	}

	// Add auth headers to request
	req.Header.Add("Authorization", bearer)

	// Send req using http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// Close the response body when finished
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result AccountNameData
	if err := json.Unmarshal(body, &result)
	err != nil {
		return "", errors.New("could not unmarshal JSON")
	}

	if result.Request.StatusCode != 200 {
		errorMsg := fmt.Sprintf("status code: %d", result.Request.StatusCode)
		return "", errors.New(errorMsg)
	}

	address := result.Response.Name

	return address, nil
}
