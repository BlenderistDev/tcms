package telegramClient

import (
	"tcms/m/internal/dry"
	"testing"
)

func TestGetMTProtoHost(t *testing.T) {
	dry.TestEnvString(t, "MTPROTO_HOST", "no mtproto host", getMTProtoHost)
}

func TestGetAppId(t *testing.T) {
	dry.TestEnvIntWithError(t, "APP_ID", "no app id", getAppId)
}
