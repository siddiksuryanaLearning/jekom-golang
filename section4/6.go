package section4

func pangkat(basis, eksponen int) int {
	result := 1
	for i := 0; i < eksponen; i++ {
		result *= basis
	}
	return result
}