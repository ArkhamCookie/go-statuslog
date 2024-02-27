package env_test

import (
	"testing"

	"internal/env"
)

var (
	err    error
	output string
	want   string
)

// Test getting a value from the env file.
func TestGetEnvValue(t *testing.T) {
	inputValue := "GET_TEST"
	inputFile := "test.env"
	want = "load-test-successful"

	output, err = env.GetEnvValue(inputValue, inputFile)
	if err != nil {
		t.Fatal("Error:", err)
	}
	if want != output {
		t.Fatalf(
			`env.GetEnvValue(%q, %q) = %q, want match for %#q`,
			inputValue, inputFile, output, want,
		)
	}
}

// Test trying to get a value that doesn't exist.
func TestGetEmptyValue(t *testing.T) {
	inputValue := "GET_FALSE"
	inputFile := "test.env"

	output, err = env.GetEnvValue(inputValue, inputFile)
	if err == nil {
		t.Fatalf(
			`env.GetEnvValue(%q, %q) = %q, wanted error to be thrown`,
			inputValue, inputFile, output,
		)
	}
}
