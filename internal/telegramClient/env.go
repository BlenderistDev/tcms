package telegramClient

import (
	"tcms/internal/dry"
)

func getTelegramBridgeHost() (string, error) {
	return dry.GetEnvStr("TELEGRAM_BRIDGE_HOST")
}
