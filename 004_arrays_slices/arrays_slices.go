package arrays_slices

import (
	"fmt"
)

func FindItemIndexInArray(haystack []int, needle int) (int, error) {
	length := len(haystack)

	if length < 3 {
		return -1, fmt.Errorf("The given haystack has a length less than 3")
	}

	for index, item := range haystack {
		if item == needle {
			return index, nil
		}
	}

	return -1, fmt.Errorf("Not found")
}

func Sum(numbers []int) (int, error) {
	length := len(numbers)

	if length < 2 {
		return -1, fmt.Errorf("The given array has a length that is less than 2 items")
	}

	result := 0
	for _, number := range numbers {
		result += number
	}

	return result, nil
}

func SumAll(pairs ...[]int) []int {
	var result []int

	for _, pair := range pairs {
		sum, err := Sum(pair)
		if err == nil {
			result = append(result, sum)
		}
	}

	return result
}
