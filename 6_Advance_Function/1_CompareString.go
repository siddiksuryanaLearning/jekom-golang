package main


func CekString(str1, str2 string) string {
    bytes1 := []byte(str1)
    bytes2 := []byte(str2)

    charCount := make(map[byte]int)
    for _, b := range bytes1 {
        charCount[b]++
    }

    var commonChars []byte
    for _, b := range bytes2 {
        if count, ok := charCount[b]; ok && count > 0 {
            commonChars = append(commonChars, b)
            charCount[b]--
        }
    }

    commonStr := string(commonChars)

    return commonStr
}

func main(){
	fmt.Println(CekString("AKA", "AKASHI"))
	fmt.Println(CekString("KANGGORO", "KANG"))
	fmt.Println(CekString("KI", "KIJANG"))
	fmt.Println(CekString("KUPU-KUPU", "KUPU"))
	fmt.Println(CekString("ILANG", "ILA"))
}