package db

import (
	"testing"

	"tcms/m/internal/dry"
)

func TestGetMongoHost(t *testing.T) {
	dry.TestEnvString(t, "MONGO_HOST", getMongoHost)
}
