package db

import (
	"os"
	"testing"
)

func TestGetMongoHost(t *testing.T) {
	host := "127.0.0.1:1111"
	err := os.Setenv("MONGO_HOST", host)
	if err != nil {
		t.Error(err)
	}
	result, err := getMongoHost()
	if err != nil {
		t.Error(err)
	}
	if result != host {
		t.Errorf("expect %s, got %s", host, result)
	}
}
