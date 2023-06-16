package section4

func isBilanganPrima(bilangan int) bool {
	if bilangan < 2 {
		return false
	}
	for i := 2; i*i <= bilangan; i++ {
		if bilangan%i == 0 {
			return false
		}
	}
	return true
}