package tcms

import (
	"tcms/m/internal/dry"
)

func getTcmsHost() (string, error) {
	return dry.GetEnvStr("TCMS_HOST")
}
