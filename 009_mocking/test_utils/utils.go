package test_utils

import "testing"

func AssertValue(t testing.TB, result string, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("Expected: %q, got: %q", expected, result)
	}
}

func AssertError(t testing.TB, err error, expected error) {
	t.Helper()
	if err == nil {
		t.Fatal("Expected to get error")
	}

	if err != expected {
		t.Errorf("Got error: %s, Expected: %s", err, expected)
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Expected NOT to get error")
	}
}
