package tcms

import (
	"tcms/internal/dry"
)

func getTcmsHost() (string, error) {
	return dry.GetEnvStr("TCMS_HOST")
}
