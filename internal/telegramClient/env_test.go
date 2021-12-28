package telegramClient

import (
	"tcms/m/internal/dry"
	"testing"
)

func TestGetMTProtoHost(t *testing.T) {
	dry.TestEnvString(t, "MTPROTO_HOST", getMTProtoHost)
}

func TestGetAppId(t *testing.T) {
	dry.TestEnvIntWithDefault(t, "TELEGRAM_APP_ID", 0, getAppId)
}

func TestGetAppHash(t *testing.T) {
	dry.TestEnvString(t, "TELEGRAM_APP_HASH", getAppHash)
}

func TestGetTelegramBridgeHost(t *testing.T) {
	dry.TestEnvString(t, "TELEGRAM_BRIDGE_HOST", getTelegramBridgeHost)
}
