package redis

import (
	"tcms/m/internal/dry"
	"testing"
)

func TestGetRedisHost(t *testing.T) {
	dry.TestEnvString(t, "REDIS_HOST", "no redis host", getRedisHost)
}

func TestGetRedisPassword(t *testing.T) {
	dry.TestEnvStringWithDefault(t, "REDIS_PASSWORD", "", getRedisPassword)
}

func TestGetRedisDatabase(t *testing.T) {
	dry.TestEnvIntWithDefault(t, "REDIS_DATABASE", 0, getRedisDatabase)
}
