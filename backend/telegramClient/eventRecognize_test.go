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

func TestParseUnknown_setInt64(t *testing.T) {
	prefix := "test"
	var inputInt64 int64 = 12
	result := "12"
	testParseUnknownSimple(t, inputInt64, prefix, result)

	var inputInt32 int32 = 12
	result = "12"
	testParseUnknownSimple(t, inputInt32, prefix, result)
}

func testParseUnknownSimple(t *testing.T, in interface{}, prefix string, parsed string) {
	result := parseUnknown(in, prefix)
	expected := map[string]string{
		prefix: parsed,
	}
	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("expect %v, got %v", expected, result)
	}
}
