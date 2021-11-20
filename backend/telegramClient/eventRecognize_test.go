package telegramClient

import (
	"reflect"
	"testing"
)

func TestGetPrefix(t *testing.T) {
	expected := "test"
	prefix, originalPrefix := getPrefix(expected)
	if prefix != "test." {
		t.Errorf("expected %s, got %s", prefix, expected)
	}
	if originalPrefix != expected {
		t.Errorf("expected %s, got %s", originalPrefix, "test")
	}
}

func TestParseUnknown_setInt32(t *testing.T) {
	var input int32 = 12
	prefix := "test"
	result := parseUnknown(input, "test")
	expected := map[string]string{
		prefix: "12",
	}
	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("expect %v, got %v", expected, result)
	}
}
