package dry

import "testing"

func TestGetEnvStr(t *testing.T) {
	const env = "SOME_ENV"
	TestEnvString(t, env, func() (string, error) {
		return GetEnvStr(env)
	})
}
