package kafka

import (
	"tcms/m/internal/dry"
)

func getKafkaHost() (string, error) {
	return dry.GetEnvStr("KAFKA_HOST")
}
