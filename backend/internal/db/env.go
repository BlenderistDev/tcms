package db

import (
	"fmt"
	"os"
)

func getMongoHost() (string, error) {
	host, exists := os.LookupEnv("MONGO_HOST")
	if !exists {
		return "", fmt.Errorf("no mongodb host")
	}
	return host, nil
}
