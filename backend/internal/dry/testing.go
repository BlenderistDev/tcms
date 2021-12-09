package dry

import (
	"os"
	"reflect"
	"testing"
)

func TestHandleError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func TestCheckEqual(t *testing.T, expected, testing interface{}) {
	if !reflect.DeepEqual(expected, testing) {
		t.Errorf("expect %v, got %v", expected, testing)
	}
}

func TestEnvString(t *testing.T, name string, error string, f func() (string, error)) {
	value := "testing"
	err := os.Setenv(name, value)
	TestHandleError(t, err)
	result, err := f()
	TestHandleError(t, err)
	TestCheckEqual(t, value, result)

	os.Clearenv()
	_, err = f()
	TestCheckEqual(t, error, err.Error())
}

func TestEnvStringWithDefault(t *testing.T, name string, def string, f func() string) {
	value := "testing"
	err := os.Setenv(name, value)
	TestHandleError(t, err)
	result := f()
	TestHandleError(t, err)
	TestCheckEqual(t, value, result)

	os.Clearenv()
	value = f()
	TestCheckEqual(t, def, value)
}
