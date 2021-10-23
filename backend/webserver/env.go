package webserver

import (
	"fmt"
	"os"
	"tcms/m/dry"
)

func getApiHost() string {
	host, exists := os.LookupEnv("API_HOST")
	if !exists {
		dry.HandleErrorPanic(fmt.Errorf("no api host provided"))
	}
	return host
}
