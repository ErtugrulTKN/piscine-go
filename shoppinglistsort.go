package piscine

func ShoppingListSort(slice []string) []string {
	var minLength int // int değişkenleri tanımladık.
	var minIndex int

	for i := range slice { // min uzunluk ve index bulduk.
		if len(slice[i]) < minLength {
			minLength = len(slice[i])
			minIndex = i
		}
	}

	for i := len(slice) - 1; i >= 0; i-- {
		if len(slice[i]) == minLength { // uzunluklar eşitse minindex ile değiştir.
			slice[i], slice[minIndex] = slice[minIndex], slice[i]
			minIndex = i
		}
	}

	return slice
}
