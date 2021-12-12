package telegramClient

import (
	"tcms/m/internal/dry"
	"testing"
)

const prefix = "test"

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

func TestParseUnknown_setInt(t *testing.T) {
	const intResult = "12"

	var inputInt64 int64 = 12
	testParseUnknownSimple(t, inputInt64, prefix, intResult)

	var inputInt32 int32 = 12
	testParseUnknownSimple(t, inputInt32, prefix, intResult)

	var inputInt = 12
	testParseUnknownSimple(t, inputInt, prefix, intResult)

	var inputNegativeInt = -12
	testParseUnknownSimple(t, inputNegativeInt, prefix, "-12")

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
}

func TestParseUnknown_setString(t *testing.T) {
	var inputString = "test"
	testParseUnknownSimple(t, inputString, prefix, "test")
}

func TestParseUnknown_setFloat(t *testing.T) {
	const floatResult = "12.12"

	var inputFloat32 float32 = 12.12
	testParseUnknownSimple(t, inputFloat32, prefix, floatResult)

	var inputFloat64 float32 = 12.12
	testParseUnknownSimple(t, inputFloat64, prefix, floatResult)
}

func TestParseUnknown_setComplex(t *testing.T) {
	const complexResult = "(1+3i)"

	var inputComplex64 complex64 = 3i + 1
	testParseUnknownSimple(t, inputComplex64, prefix, complexResult)

	var inputComplex128 complex128 = 3i + 1
	testParseUnknownSimple(t, inputComplex128, prefix, complexResult)
}

func TestParseUnknown_setBool(t *testing.T) {
	testParseUnknownSimple(t, true, prefix, "true")
	testParseUnknownSimple(t, false, prefix, "false")
}

func TestParseUnknown_setSlice(t *testing.T) {
	inputSlice := []int{11, 22, 33}
	result := parseUnknown(inputSlice, prefix)
	expected := map[string]string{
		prefix + "." + "0": "11",
		prefix + "." + "1": "22",
		prefix + "." + "2": "33",
	}
	dry.TestCheckEqual(t, expected, result)
}

func TestParseUnknown_setMap(t *testing.T) {
	inputMap := map[string]int{
		"test1": 1,
		"test2": 2,
		"test3": 3,
	}
	result := parseUnknown(inputMap, prefix)
	expected := map[string]string{
		prefix + "." + "test1": "1",
		prefix + "." + "test2": "2",
		prefix + "." + "test3": "3",
	}
	dry.TestCheckEqual(t, expected, result)
}

func TestParseUnknown_setStruct(t *testing.T) {
	type testStruct struct {
		Data      string
		OtherData int
	}
	inputStruct := testStruct{
		Data:      "test",
		OtherData: 1,
	}
	result := parseUnknown(inputStruct, prefix)
	expected := map[string]string{
		prefix + "." + "Data":      "test",
		prefix + "." + "OtherData": "1",
	}
	dry.TestCheckEqual(t, expected, result)
}

func testParseUnknownSimple(t *testing.T, in interface{}, prefix string, parsed string) {
	result := parseUnknown(in, prefix)
	expected := map[string]string{
		prefix: parsed,
	}
	dry.TestCheckEqual(t, expected, result)
}
