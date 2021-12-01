package dry

import "testing"

func TestHandleError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
