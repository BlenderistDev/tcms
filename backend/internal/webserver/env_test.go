package webserver

import (
	"tcms/m/internal/dry"
	"testing"
)

func TestGetApiHost(t *testing.T) {
	dry.TestEnvString(t, "API_HOST", getApiHost)
}
