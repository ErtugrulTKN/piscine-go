package piscine

func StringToIntSlice(str string) []int {
	if str == "" { // Boşsa değer yok
		return nil
	}

	result := []int{}

	for _, char := range str { // ASCII karakterlerin değerlerini içerenleri döndür.
		result = append(result, int(char)) // append result ile int dizisini birleştir.
	}

	return result
}
