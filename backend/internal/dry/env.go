package dry

import (
	"fmt"
	"os"
	"strconv"
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

func GetEnvInt(key string) (int, error) {
	str, exists := os.LookupEnv(key)
	var value int
	var err error
	if !exists {
		value = 0
	} else {
		value, err = strconv.Atoi(str)
	}
	return value, err
}
