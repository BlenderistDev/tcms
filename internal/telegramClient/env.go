package telegramClient

import (
	"tcms/m/internal/dry"
)

func getMTProtoHost() (string, error) {
	return dry.GetEnvStr("MTPROTO_HOST")
}

func getAppId() (int, error) {
	return dry.GetEnvIntWithDefault("TELEGRAM_APP_ID", 0)
}

func getAppHash() (string, error) {
	return dry.GetEnvStr("TELEGRAM_APP_HASH")
}
