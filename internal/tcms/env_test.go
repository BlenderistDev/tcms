package tcms

import (
	"testing"

	"tcms/internal/dry"
)

func TestTcmsHost(t *testing.T) {
	dry.TestEnvString(t, "TCMS_HOST", getTcmsHost)
}
