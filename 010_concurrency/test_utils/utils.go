package test_utils

import (
	"reflect"
	"testing"
)

func AssertValueDeepEqual(t testing.TB, expected any, result any) {
	t.Helper()
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}
