package telegramClient

import (
	"tcms/m/internal/dry"
)

func getTelegramBridgeHost() (string, error) {
	return dry.GetEnvStr("TELEGRAM_BRIDGE_HOST")
}
