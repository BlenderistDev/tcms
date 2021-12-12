package webserver

import (
	"tcms/m/internal/dry"
)

func getApiHost() (string, error) {
	return dry.GetEnvStr("API_HOST")
}
