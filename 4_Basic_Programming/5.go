package section4

import ( 
	"strings"
)

func isPalindrome(str string) bool {
	// Menghapus spasi dan mengubah huruf menjadi lowercase
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ToLower(str)

	// Menentukan indeks awal dan indeks akhir
	i := 0
	j := len(str) - 1

	// Memeriksa apakah karakter di indeks i sama dengan karakter di indeks j
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}