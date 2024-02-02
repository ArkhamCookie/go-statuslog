package env

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Get a value for env
func GetEnvValue(target string, file string) (string, error) {
	// Set file to '.env' if no file is given.
	// Makes file value optional.
	if file == "" {
		file = ".env"
	}

	// Load env file
	err := godotenv.Load(file)
	// Check for errors in loading env file
	if err != nil {
		log.Fatalln("Some error occurred. Error:", err)
	}

	// Get target value
	value := os.Getenv(target)
	// If value isn't defined, throw an error.
	if value == "" {
		log.Printf("%s value not found\n", target)
		return "", errors.New("value not found")
	}
	return value, nil
}
