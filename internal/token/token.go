package token

import "os"

func SetTokenEnv(token string) {
	os.Setenv("OMGLOL_TOKEN", token)
}

func GetTokenEnv() string {
	return os.Getenv("OMGLOL_TOKEN")
}
