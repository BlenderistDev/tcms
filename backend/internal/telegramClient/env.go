package telegramClient

import (
	"fmt"
	"os"
	"strconv"
	"tcms/m/internal/dry"
)

func getMTProtoHost() (string, error) {
	return dry.GetEnvStr("MTPROTO_HOST")
}

func getAppId() (int, error) {
	appId, exists := os.LookupEnv("TELEGRAM_APP_ID")
	if !exists {
		return 0, fmt.Errorf("no app key")
	}
	return strconv.Atoi(appId)
}

func getAppHash() (string, error) {
	return dry.GetEnvStr("TELEGRAM_APP_HASH")
}
