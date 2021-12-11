package dry

import (
	"fmt"
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

func TestEnvString(t *testing.T, name string, f func() (string, error)) {
	value := "testing"
	err := os.Setenv(name, value)
	TestHandleError(t, err)
	result, err := f()
	TestHandleError(t, err)
	TestCheckEqual(t, value, result)

	os.Clearenv()
	_, err = f()
	TestCheckEqual(t, fmt.Sprintf("no %s env", name), err.Error())
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

func TestEnvIntWithDefault(t *testing.T, name string, def int, f func() (int, error)) {
	value := 3
	valueStr := "3"
	err := os.Setenv(name, valueStr)
	TestHandleError(t, err)
	result, err := f()
	TestHandleError(t, err)
	TestCheckEqual(t, value, result)

	os.Clearenv()
	result, err = f()
	TestHandleError(t, err)
	TestCheckEqual(t, def, result)

	os.Clearenv()
	value, err = f()
	TestHandleError(t, err)
	TestCheckEqual(t, def, value)
}
