package tcms

import (
	"tcms/m/internal/dry"
	"testing"
)

func TestTcmsHost(t *testing.T) {
	dry.TestEnvString(t, "TCMS_HOST", getTcmsHost)
}
