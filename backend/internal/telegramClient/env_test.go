package telegramClient

import (
	"os"
	"tcms/m/internal/dry"
	"testing"
)

func TestGetMTProtoHost(t *testing.T) {
	host := "127.0.0.1:1111"
	err := os.Setenv("MTPROTO_HOST", host)
	dry.TestHandleError(t, err)
	result, err := getMTProtoHost()
	dry.TestHandleError(t, err)
	dry.TestCheckEqual(t, host, result)
}

func TestGetMTProtoHost_notExist(t *testing.T) {
	os.Clearenv()
	_, err := getMTProtoHost()
	dry.TestCheckEqual(t, "no mtproto host", err.Error())
}

func TestGetAppId(t *testing.T) {
	appId := "123123"
	appIdInt := 123123
	err := os.Setenv("TELEGRAM_APP_ID", appId)
	dry.TestHandleError(t, err)
	result, err := getAppId()
	dry.TestHandleError(t, err)
	dry.TestCheckEqual(t, appIdInt, result)
}
