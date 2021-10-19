package webserver

import (
	"fmt"
	"os"
)

func getApiHost() (string, error) {
	host, exists := os.LookupEnv("API_HOST")
	if !exists {
		return "", fmt.Errorf("no api host provided")
	}
	return host, nil
}
