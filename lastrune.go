package piscine

func LastRune(s string) rune {
	if len(s) == 0 {
		return 0
	}
	runes := []rune(s)
	return runes[len(runes)-1]
}
