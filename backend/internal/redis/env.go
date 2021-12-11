package redis

import (
	"tcms/m/internal/dry"
)

func getRedisHost() (string, error) {
	return dry.GetEnvStr("REDIS_HOST")
}

func getRedisPassword() string {
	return dry.GetEnvStrWithDefault("REDIS_PASSWORD", "")
}

func getRedisDatabase() (int, error) {
	return dry.GetEnvInt("REDIS_DATABASE")
}
