package piscine

func JumpOver(str string) string {
	if len(str) == 0 {
		return "\n"
	}

	var output string
	for i := 2; i < len(str); i += 3 { // Üçüncü karakterden (indeks 2) başlayarak giriş dizesindeki karakterleri yinele.
		output += string(str[i]) // Her üç karakterden birini dizeye ekleyin.
	}

	return output + "\n"
}
