package db

import (
	"os"
	"tcms/m/internal/dry"
	"testing"
)

func TestGetMongoHost(t *testing.T) {
	host := "127.0.0.1:1111"
	err := os.Setenv("MONGO_HOST", host)
	dry.TestHandleError(t, err)
	result, err := getMongoHost()
	dry.TestHandleError(t, err)
	if result != host {
		t.Errorf("expect %s, got %s", host, result)
	}
}

func TestGetMongoHost_notExist(t *testing.T) {
	os.Clearenv()
	_, err := getMongoHost()
	dry.TestCheckEqual(t, "no mongodb host", err.Error())
}
