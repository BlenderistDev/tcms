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
	const prefix = "test"
	const intResult = "12"
	const floatResult = "12.12"
	const complexResult = "(1+3i)"

	var inputInt64 int64 = 12
	testParseUnknownSimple(t, inputInt64, prefix, intResult)

	var inputInt32 int32 = 12
	testParseUnknownSimple(t, inputInt32, prefix, intResult)

	var inputInt = 12
	testParseUnknownSimple(t, inputInt, prefix, intResult)

	var inputInt8 int8 = 12
	testParseUnknownSimple(t, inputInt8, prefix, intResult)

	var inputInt16 int16 = 12
	testParseUnknownSimple(t, inputInt16, prefix, intResult)

	var inputUint uint = 12
	testParseUnknownSimple(t, inputUint, prefix, intResult)

	var inputUint8 uint8 = 12
	testParseUnknownSimple(t, inputUint8, prefix, intResult)

	var inputUint16 uint16 = 12
	testParseUnknownSimple(t, inputUint16, prefix, intResult)

	var inputUint32 uint32 = 12
	testParseUnknownSimple(t, inputUint32, prefix, intResult)

	var inputUint64 uint32 = 12
	testParseUnknownSimple(t, inputUint64, prefix, intResult)

	var inputString = "test"
	testParseUnknownSimple(t, inputString, prefix, "test")

	var inputFloat32 float32 = 12.12
	testParseUnknownSimple(t, inputFloat32, prefix, floatResult)

	var inputFloat64 float32 = 12.12
	testParseUnknownSimple(t, inputFloat64, prefix, floatResult)

	var inputComplex64 complex64 = 3i + 1
	testParseUnknownSimple(t, inputComplex64, prefix, complexResult)

	var inputComplex128 complex128 = 3i + 1
	testParseUnknownSimple(t, inputComplex128, prefix, complexResult)

	testParseUnknownSimple(t, true, prefix, "true")
	testParseUnknownSimple(t, false, prefix, "false")
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
