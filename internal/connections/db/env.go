package db

import (
	"tcms/m/internal/dry"
)

func getMongoHost() (string, error) {
	return dry.GetEnvStr("MONGO_HOST")
}
