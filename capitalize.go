package piscine

func Capitalize(s string) string {
	Runes := []rune(s)
	for index, letter := range Runes {
		if checkAlphNum(letter) {
			if index == 0 || checkAlphNum(Runes[index-1]) == false {
				if letter >= 'a' && letter <= 'z' {
					Runes[index] = letter - 32
				}
			} else {
				if letter >= 'A' && letter <= 'Z' {
					Runes[index] = letter + 32
				}
			}
		}
	}
	return string(Runes)
}

func checkAlphNum(r rune) bool {
	if r >= 'A' && r <= 'Z' ||
		r >= 'a' && r <= 'z' ||
		r >= '0' && r <= '9' {
		return true
	}
	return false
}
