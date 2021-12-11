package dry

import "testing"

func TestGetEnvStr(t *testing.T) {
	const env = "SOME_ENV"
	TestEnvString(t, env, func() (string, error) {
		return GetEnvStr(env)
	})
}

func TestGetEnvStrWithDefault(t *testing.T) {
	const (
		env = "SOME_ENV"
		def = "default"
	)
	TestEnvStringWithDefault(t, env, def, func() string {
		return GetEnvStrWithDefault(env, def)
	})
}

func TestGetEnvIntWithDefault(t *testing.T) {
	const (
		env = "SOME_ENV"
		def = 3
	)
	TestEnvIntWithDefault(t, env, def, func() (int, error) {
		return GetEnvIntWithDefault(env, def)
	})
}
