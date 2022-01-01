package telegramClient

import (
	"tcms/m/internal/dry"
	"testing"
)

func TestGetTelegramBridgeHost(t *testing.T) {
	dry.TestEnvString(t, "TELEGRAM_BRIDGE_HOST", getTelegramBridgeHost)
}
