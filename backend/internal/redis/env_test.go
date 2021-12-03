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
