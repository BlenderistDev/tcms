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
