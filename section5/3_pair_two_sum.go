package section5

import "fmt"

func findPair(array []int, target int) (int, int) {
	left := 0
	right := len(array) - 1

	for left < right {
		sum := array[left] + array[right]

		if sum == target {
			return left, right
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return -1, -1 // If no pair is found
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 10

	index1, index2 := findPair(array, target)
	if index1 != -1 && index2 != -1 {
		fmt.Printf("Indices of the pair: %d, %d\n", index1, index2)
	} else {
		fmt.Println("No pair found.")
	}
}
