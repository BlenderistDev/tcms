package db

import (
	"tcms/internal/dry"
)

func getMongoHost() (string, error) {
	return dry.GetEnvStr("MONGO_HOST")
}
