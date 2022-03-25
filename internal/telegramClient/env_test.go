package telegramClient

import (
	"testing"

	"tcms/m/internal/dry"
)

func TestGetTelegramBridgeHost(t *testing.T) {
	dry.TestEnvString(t, "TELEGRAM_BRIDGE_HOST", getTelegramBridgeHost)
}
