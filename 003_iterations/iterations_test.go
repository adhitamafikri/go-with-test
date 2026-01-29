package iterations

import (
	"testing"
)

func TestIteration(t *testing.T) {
	t.Run("RepeatAndConcat() function properly produce a string that is formed by concatenating the input string n times", func(t *testing.T) {
		result, _ := RepeatAndConcat("cat", 10)
		expected := "catcatcatcatcatcatcatcatcatcat"

		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("RepeatAndConcat() function would raise an error if the given n is less than 3", func(t *testing.T) {
		_, err := RepeatAndConcat("cat", 1)

		if err == nil {
			t.Errorf("Expected to throw error")
		}
	})
}

func BenchmarkRepeatAndConcat(b *testing.B) {
	for b.Loop() {
		RepeatAndConcat("Sane", 15)
	}
}
