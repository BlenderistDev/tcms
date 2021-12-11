package telegramClient

import (
	"tcms/m/internal/dry"
	"testing"
)

func TestGetMTProtoHost(t *testing.T) {
	dry.TestEnvString(t, "MTPROTO_HOST", "no MTPROTO_HOST env", getMTProtoHost)
}

func TestGetAppId(t *testing.T) {
	dry.TestEnvIntWithError(t, "TELEGRAM_APP_ID", "no app key", getAppId)
}

func TestGetAppHash(t *testing.T) {
	dry.TestEnvString(t, "TELEGRAM_APP_HASH", "no TELEGRAM_APP_HASH env", getAppHash)
}
