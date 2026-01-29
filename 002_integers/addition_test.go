package integers

import (
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	t.Run("Should properly add two integers", func(t *testing.T) {
		sum := Add(15, 25)
		expected := 40

		if sum != expected {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	})
}

func ExampleAdd() {
	sum := Add(12, 17)
	fmt.Println(sum)
	// Output: 29
}
