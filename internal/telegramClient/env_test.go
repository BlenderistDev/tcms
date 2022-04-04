package telegramClient

import (
	"testing"

	"tcms/internal/dry"
)

func TestGetTelegramBridgeHost(t *testing.T) {
	dry.TestEnvString(t, "TELEGRAM_BRIDGE_HOST", getTelegramBridgeHost)
}
