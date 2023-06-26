package main

import "fmt"

func findMinMax(numbers ...*int) (min int, max int) {
	if len(numbers) == 0 {
		return
	}

	min = *numbers[0]
	max = *numbers[0]

	for _, num := range numbers {
		if *num < min {
			min = *num
		}
		if *num > max {
			max = *num
		}
	}

	return min, max
}

func main() {
	a := 5
	b := 3
	c := 8
	d := 2
	e := 10
	f := 1

	min, max := findMinMax(&a, &b, &c, &d, &e, &f)

	fmt.Println("Minimum:", min)
	fmt.Println("Maximum:", max)
}
