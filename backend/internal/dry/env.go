package dry

import (
	"fmt"
	"os"
)

func GetEnvStr(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("no %s env", key)
	}
	return value, nil
}

func GetEnvStrWithDefault(key string, def string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = def
	}
	return value
}
