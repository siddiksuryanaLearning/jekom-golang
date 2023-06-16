package main

import (
	"fmt"
<<<<<<< HEAD
	"jekom-golang/section4"
	"jekom-golang/section4/2"
	"jekom-golang/section4/3"
	"jekom-golang/section4/4"
	"jekom-golang/section4/5"
	"jekom-golang/section4/6"
	"jekom-golang/section4/7"
	"jekom-golang/section4/8"
=======
	"/section4/1.go"
	"/section4/2.go"
	"/section4/3.go"
	"/section4/4.go"
	"/section4/5.go"
	"/section4/6.go"
	"/section4/7.go"
	"/section4/8.go"
>>>>>>> origin

	"github.com/siddiksuryanaLearning/jekom-golang"
)

func main() {
	
	fmt.Printf("==== No.1 ====")
	jariJari := 3.0
	tinggi := 5.0
	luasPermukaan := hitungLuasPermukaanTabung(jariJari, tinggi)
	fmt.Printf("Luas Permukaan Tabung: %.2f\n", luasPermukaan)

	fmt.Printf("=== No.2 ===")
	nilai := 80
	grade := konversiGrade(nilai)
	fmt.Printf("Nilai: %d, Grade: %s\n", nilai, grade)

	fmt.Printf("=== No.3 ===")
	bilangan := 6	
	faktor := faktorBilangan(bilangan)
	fmt.Printf("Faktor-faktor dari bilangan %d adalah: %v\n", bilangan, faktor)

	fmt.Printf("=== No.4 ===")
	number := 17
	if isBilanganPrima(number) {
		fmt.Printf("%d adalah bilangan prima.\n", number)
	} else {
		fmt.Printf("%d bukan bilangan prima.\n", number)
	}

	fmt.Printf("=== No.5 ===")
	kata := "Kasur ini rusak"
	if isPalindrome(kata) {
		fmt.Printf("%s adalah palindrome.\n", kata)
	} else {
		fmt.Printf("%s bukan palindrome.\n", kata)
	}

	fmt.Printf("=== No.6 ===")
	basis := 2
	eksponen := 3
	hasil := pangkat(basis, eksponen)
	fmt.Printf("%d pangkat %d = %d\n", basis, eksponen, hasil)

	fmt.Printf("=== No.7 ===")
	tinggi := 5
	cetakSegitigaAsterisk(tinggi)

	fmt.Printf("=== No.8 ===")
	tinggi := 5
	cetakPerkalian(tinggi)

}
