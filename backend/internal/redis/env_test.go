package redis

import (
	"os"
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
