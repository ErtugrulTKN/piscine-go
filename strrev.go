package piscine

func StrRev(s string) string {
	runes := []rune(s)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-i-1] = runes[length-i-1], runes[i]
	}

	return string(runes)
}
