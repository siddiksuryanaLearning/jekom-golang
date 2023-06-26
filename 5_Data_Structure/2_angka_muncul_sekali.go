package main

import "fmt"

func removeDuplicates(arr []int) []int {
	unique := make(map[int]bool)
	uniqueArray := []int{}

	for _, num := range arr {
		if !unique[num] {
			unique[num] = true
			uniqueArray = append(uniqueArray, num)
		}
	}

	return uniqueArray
}

func main() {
	array := []int{7,6,5,2,3,7,5,2}
	uniqueArray := removeDuplicates(array)
	fmt.Println("Unique Array:", uniqueArray)
}
