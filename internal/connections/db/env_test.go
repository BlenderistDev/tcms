package db

import (
	"testing"

	"tcms/internal/dry"
)

func TestGetMongoHost(t *testing.T) {
	dry.TestEnvString(t, "MONGO_HOST", getMongoHost)
}
