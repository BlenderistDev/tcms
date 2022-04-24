package dry

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvString(t *testing.T, name string, f func() (string, error)) {
	value := "testing"
	err := os.Setenv(name, value)
	assert.Nil(t, err)
	result, err := f()
	assert.Nil(t, err)
	assert.Equal(t, value, result)

	os.Clearenv()
	_, err = f()
	assert.Equal(t, fmt.Sprintf("no %s env", name), err.Error())
}

func TestEnvStringWithDefault(t *testing.T, name string, def string, f func() string) {
	value := "testing"
	err := os.Setenv(name, value)
	assert.Nil(t, err)
	result := f()
	assert.Nil(t, err)
	assert.Equal(t, value, result)

	os.Clearenv()
	value = f()
	assert.Equal(t, def, value)
}

func TestEnvIntWithDefault(t *testing.T, name string, def int, f func() (int, error)) {
	value := 3
	valueStr := "3"
	err := os.Setenv(name, valueStr)
	assert.Nil(t, err)
	result, err := f()
	assert.Nil(t, err)
	assert.Equal(t, value, result)

	os.Clearenv()
	result, err = f()
	assert.Nil(t, err)
	assert.Equal(t, def, result)

	os.Clearenv()
	value, err = f()
	assert.Nil(t, err)
	assert.Equal(t, def, value)
}
