package piscine

func Unmatch(a []int) int {
	counts := make(map[int]int) // int saklamak için map oluşturduk.

	for _, MyNum := range a {
		counts[MyNum]++
	}

	for _, MyNum := range a { // eşleştirilmemiş ögeyi kontrol et.
		if counts[MyNum]%2 == 1 { // sayının çiftleri var ise return -1
			return MyNum
		}
	}

	return -1
}
