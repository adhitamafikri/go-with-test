package test_utils

import (
	"reflect"
	"testing"
)

func AssertValue(t testing.TB, result interface{}, expected interface{}) {
	t.Helper()
	if result != expected {
		t.Errorf("Expected: %q, got: %q", expected, result)
	}
}

func AssertValueDeepEqual(t testing.TB, result interface{}, expected interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}

func AssertCancelled(t testing.TB, result bool) {
	t.Helper()
	if !result {
		t.Error("context: It should be cancelled")
	}
}

func AssertNotCancelled(t testing.TB, result bool) {
	t.Helper()
	if result {
		t.Error("context: It should NOT be cancelled")
	}
}

func AssertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, item := range haystack {
		if item == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
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
