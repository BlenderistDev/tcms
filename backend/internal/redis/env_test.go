package redis

import (
	"os"
	"strconv"
	"tcms/m/internal/dry"
	"testing"
)

func TestGetRedisHost(t *testing.T) {
	host := "127.0.0.1:1111"
	err := os.Setenv("REDIS_HOST", host)
	dry.TestHandleError(t, err)
	result, err := getRedisHost()
	dry.TestHandleError(t, err)
	dry.TestCheckEqual(t, host, result)
}

func TestGetRedisHost_notExist(t *testing.T) {
	os.Clearenv()
	_, err := getRedisHost()
	dry.TestCheckEqual(t, "no redis host", err.Error())
}

func TestGetRedisPassword(t *testing.T) {
	password := "pass"
	err := os.Setenv("REDIS_PASSWORD", password)
	dry.TestHandleError(t, err)
	result := getRedisPassword()
	dry.TestHandleError(t, err)
	dry.TestCheckEqual(t, password, result)
}

func TestGetRedisPasswordEmpty(t *testing.T) {
	os.Clearenv()
	result := getRedisPassword()
	dry.TestCheckEqual(t, "", result)
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
