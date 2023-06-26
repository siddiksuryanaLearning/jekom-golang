package section4

import "fmt"

func cetakPerkalian(tinggi int) {
	for i := 1; i <= tinggi; i++ {
		for j := 1; j <= tinggi; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Println()
	}
}

