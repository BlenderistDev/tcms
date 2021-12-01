package dry

import (
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
