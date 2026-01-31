package test_utils

import (
	"reflect"
	"testing"
)

func AssertValue(t testing.TB, result string, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("Expected: %q, got: %q", expected, result)
	}
}

func AssertValueDeepEqual(t testing.TB, expected any, result any) {
	t.Helper()
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected: %v, got: %v", expected, result)
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
