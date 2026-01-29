package integers

import "fmt"

func Add(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	fmt.Printf("Result of Add(25 + 60) = %d", Add(25, 60))
}
