package action

import (
	"github.com/joho/godotenv"
	"os"
	"testing"
)

// for correct config files loading
func TestMain(m *testing.M) {
	err := os.Chdir("../../")
	if err != nil {
		panic(err)
	}
	err = godotenv.Load()
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
