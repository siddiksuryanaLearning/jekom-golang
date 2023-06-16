package section4

import "fmt"

func cetakSegitigaAsterisk(tinggi int) {
	for i := 1; i <= tinggi; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}


	
