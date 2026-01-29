package arrays_slices

import (
	"slices"
	"testing"
)

func TestFindItemIndexInArray(t *testing.T) {
	t.Run("FindItemIndexInArray() will return the index of the first occurrence of the needle", func(t *testing.T) {
		result, _ := FindItemIndexInArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, 10)
		expected := 9

		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("FindItemIndexInArray() will return an error if it does not found the needle", func(t *testing.T) {
		_, err := FindItemIndexInArray([]int{100, 250, 350}, 259)

		if err == nil {
			t.Errorf("Expected to get error, cannot find the needle")
		}
	})

	t.Run("FindItemIndexInArray() will return an error if the given haystack has a length less than 3", func(t *testing.T) {
		_, err := FindItemIndexInArray([]int{100, 250}, 10)

		if err == nil {
			t.Errorf("Expected to get error, given haystack length is less than 3")
		}
	})
}

func BenchmarkFindItemIndexInArray(b *testing.B) {
	for b.Loop() {
		FindItemIndexInArray([]int{20, 150, 230, 1125, 2150}, 1125)
	}
}

func TestSum(t *testing.T) {
	t.Run("Sum() will return a single integer which is a result of the SUM operations of the numbers in the given array", func(t *testing.T) {
		result, err := Sum([]int{10, 25, 39, 75})
		expected := 149

		if err != nil {
			t.Errorf("Error should not be raised, array has length > 2")
		}
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Sum() raise an error if the given array has less than 2 items", func(t *testing.T) {
		result, err := Sum([]int{10})
		expected := -1

		if err == nil {
			t.Errorf("Error should be raised, array has length < 2")
		}
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("SumAll() will return an array of results produced by the sum operations of each arrays supplied", func(t *testing.T) {
		result := SumAll([]int{1, 2, 3, 4}, []int{15, 25}, []int{25000, 34500})
		expected := []int{10, 40, 59500}

		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
