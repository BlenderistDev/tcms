package kafka

import (
	"tcms/m/internal/dry"
	"testing"
)

func TestGetKafkaHost(t *testing.T) {
	dry.TestEnvString(t, "KAFKA_HOST", getKafkaHost)
}
