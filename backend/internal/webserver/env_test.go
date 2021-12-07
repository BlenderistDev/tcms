package webserver

import (
	"os"
	"tcms/m/internal/dry"
	"testing"
)

func TestGetApiHost(t *testing.T) {
	host := "127.0.0.1:1111"
	err := os.Setenv("API_HOST", host)
	dry.TestHandleError(t, err)
	result := getApiHost()
	dry.TestHandleError(t, err)
	dry.TestCheckEqual(t, host, result)
}
