package helpers

import (
	"os"
)

func GoDotEnvVar(key string) string {

	// load .env file
	return os.Getenv(key)
}
