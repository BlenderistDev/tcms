package telegramClient

import (
	"fmt"
	"os"
	"strconv"
)

func getMTProtoHost() (string, error) {
	host, exists := os.LookupEnv("MTPROTO_HOST")
	if !exists {
		return "", fmt.Errorf("no mtproto host")
	}
	return host, nil
}

func getAppId() (int, error) {
	appId, exists := os.LookupEnv("TELEGRAM_APP_ID")
	if !exists {
		return 0, fmt.Errorf("no app key")
	}
	return strconv.Atoi(appId)
}

func getAppHash() (string, error) {
	appHash, exists := os.LookupEnv("TELEGRAM_APP_HASH")
	if !exists {
		return "", fmt.Errorf("no app hash")
	}
	return appHash, nil
}
