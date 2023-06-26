package section5

import "fmt"

func mergeArrays(arr1 []string, arr2 []string) []string {
	merged := make(map[string]bool)
	mergedArray := []string{}

	for _, str := range arr1 {
		merged[str] = true
	}

	for _, str := range arr2 {
		if !merged[str] {
			merged[str] = true
		}
	}

	for str := range merged {
		mergedArray = append(mergedArray, str)
	}

	return mergedArray
}

func main() {
	array1 := []string{"lee", "jin"}
	array2 := []string{"kazuya", "panda"}
	mergedArray := mergeArrays(array1, array2)
	fmt.Println("Merged Array:", mergedArray)
}
