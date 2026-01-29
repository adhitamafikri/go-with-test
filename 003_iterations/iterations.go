package iterations

import (
	"fmt"
	"strings"
)

func RepeatAndConcat(str string, n int) (string, error) {
	if n < 3 {
		return "", fmt.Errorf("The n could not be less than 3.\nReceived: %d", n)
	}

	// result := ""
	// for i := 0; i < n; i++ {
	// 	result += str
	// }

	// More memory efficient approach using strings.Builder
	// This avoid copying memory to accommodate new string
	var result strings.Builder
	for i := 0; i < n; i++ {
		result.WriteString(str)
	}

	return result.String(), nil
}
