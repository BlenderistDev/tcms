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

	var inputInt = 12
	result = "12"
	testParseUnknownSimple(t, inputInt, prefix, result)

	var inputInt8 int8 = 12
	result = "12"
	testParseUnknownSimple(t, inputInt8, prefix, result)

	var inputInt16 int16 = 12
	result = "12"
	testParseUnknownSimple(t, inputInt16, prefix, result)

	var inputString = "test"
	result = "test"
	testParseUnknownSimple(t, inputString, prefix, result)

	var inputFloat32 float32 = 12.12
	result = "12.12"
	testParseUnknownSimple(t, inputFloat32, prefix, result)

	var inputFloat64 float32 = 12.12
	result = "12.12"
	testParseUnknownSimple(t, inputFloat64, prefix, result)
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
