package redis

import (
	"fmt"
	"os"
	"strconv"
)

func getRedisHost() (string, error) {
	host, exists := os.LookupEnv("REDIS_HOST")
	if !exists {
		return "", fmt.Errorf("no redis host")
	}
	return host, nil
}

func getRedisPassword() string {
	password, exists := os.LookupEnv("REDIS_PASSWORD")
	if !exists {
		password = ""
	}
	return password
}

func getRedisDatabase() (int, error) {
	str, exists := os.LookupEnv("REDIS_DATABASE")
	var database int
	var err error
	if !exists {
		database = 0
	} else {
		database, err = strconv.Atoi(str)
	}
	return database, err
}
