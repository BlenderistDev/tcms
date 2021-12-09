package telegramClient

import (
	"tcms/m/internal/dry"
	"testing"
)

func TestGetMTProtoHost(t *testing.T) {
	dry.TestEnvString(t, "MTPROTO_HOST", "no mtproto host", getMTProtoHost)
}
