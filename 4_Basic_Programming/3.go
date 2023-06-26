package section4

func faktorBilangan(bilangan int) []int {
	faktor := []int{}
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			faktor = append(faktor, i)
		}
	}
	return faktor
}
