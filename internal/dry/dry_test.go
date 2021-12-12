package dry

import (
	"fmt"
	"testing"
)

func TestHandleErrorPanic(t *testing.T) {
	HandleErrorPanic(nil)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expect panic")
		}
	}()
	HandleErrorPanic(fmt.Errorf("some error"))
}
