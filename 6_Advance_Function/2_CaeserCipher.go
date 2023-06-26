package main

import "fmt"

func CaesarCipher(shift int, text string) string {
	shift = shift % 26
	if shift < 0 {
		shift += 26
	}
	bytes := []byte(text)
	for i, b := range bytes {
		if b >= 'A' && b <= 'Z' {
			bytes[i] = (b-'A'+byte(shift))%26 + 'A'
		} else if b >= 'a' && b <= 'z' {
			bytes[i] = (b-'a'+byte(shift))%26 + 'a'
		}
	}
	cipherText := string(bytes)
	return cipherText
}

func main(){
	fmt.Println(CaesarCipher(3, "abc"))
	fmt.Println(CaesarCipher(2, "alta"))
	fmt.Println(CaesarCipher(10, "alterraacademy"))
	fmt.Println(CaesarCipher(1, "abcdefghijklmnoprstuvwxyz"))
	fmt.Println(CaesarCipher(1000, "abcdefghijklmnopqrstuvwxyz"))

}

