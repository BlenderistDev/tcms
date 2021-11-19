package telegramClient

import "testing"

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
