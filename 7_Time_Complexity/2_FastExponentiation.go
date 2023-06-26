package main

import "fmt"

func power(x, n int) int {
	result := 1

	for i := 0; i < n; i++ {
		result *= x
	}

	return result
}

func main() {
	x := 2
	n := 5

	result := power(x, n)
	fmt.Printf("%d^%d = %d\n", x, n, result)
}
