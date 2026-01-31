package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Adhitama")

	result := buffer.String()
	expected := "Hello, Adhitama"

	assertValue(t, result, expected)
}

func assertValue(t testing.TB, result string, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("Expected: %q, got: %q", expected, result)
	}
}
