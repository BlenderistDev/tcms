package action

import (
	"os"
	"tcms/m/internal/dry"
	"testing"
)

// TestMain for correct config data loading
func TestMain(m *testing.M) {
	err := os.Setenv("TELEGRAM_APP_HASH", "test")
	dry.HandleErrorPanic(err)

	err = os.Setenv("MTPROTO_HOST", "test")
	dry.HandleErrorPanic(err)

	f, err := os.Create("tg_public_keys.pem")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			dry.HandleErrorPanic(err)
		}
	}(f)
	dry.HandleErrorPanic(err)

	_, err = f.WriteString("-----BEGIN RSA PUBLIC KEY-----\n")
	dry.HandleErrorPanic(err)
	_, err = f.WriteString("test\n")
	dry.HandleErrorPanic(err)
	_, err = f.WriteString("-----END RSA PUBLIC KEY-----\n")
	dry.HandleErrorPanic(err)

	m.Run()

	err = os.Remove("tg_public_keys.pem")
	dry.HandleErrorPanic(err)
}
