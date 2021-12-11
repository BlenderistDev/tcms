package redis

import (
	"os"
	"strconv"
	"tcms/m/internal/dry"
)

func getRedisHost() (string, error) {
	return dry.GetEnvStr("REDIS_HOST")
}

func getRedisPassword() string {
	return dry.GetEnvStrWithDefault("REDIS_PASSWORD", "")
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
