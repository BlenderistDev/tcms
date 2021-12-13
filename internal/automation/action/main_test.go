package action

import (
	"github.com/joho/godotenv"
	"os"
	"tcms/m/internal/dry"
	"testing"
)

// for correct config files loading
func TestMain(m *testing.M) {
	err := os.Chdir("../../")
	dry.HandleErrorPanic(err)
	err = godotenv.Load()
	dry.HandleErrorPanic(err)
	os.Exit(m.Run())
}
