package redis

import (
	"os"
	"strconv"
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
	database := 1
	databaseStr := strconv.Itoa(database)
	err := os.Setenv("REDIS_DATABASE", databaseStr)
	dry.TestHandleError(t, err)
	result, err := getRedisDatabase()
	dry.TestHandleError(t, err)
	dry.TestCheckEqual(t, database, result)
}

func TestGetRedisDatabaseEmpty(t *testing.T) {
	os.Clearenv()
	result, err := getRedisDatabase()
	dry.TestHandleError(t, err)
	dry.TestCheckEqual(t, 0, result)
}
