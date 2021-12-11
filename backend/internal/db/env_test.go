package db

import (
	"tcms/m/internal/dry"
	"testing"
)

func TestGetMongoHost(t *testing.T) {
	dry.TestEnvString(t, "MONGO_HOST", getMongoHost)
}
