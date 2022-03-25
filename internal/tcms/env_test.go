package tcms

import (
	"testing"

	"tcms/m/internal/dry"
)

func TestTcmsHost(t *testing.T) {
	dry.TestEnvString(t, "TCMS_HOST", getTcmsHost)
}
