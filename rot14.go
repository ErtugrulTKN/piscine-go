package piscine

func Rot14(s string) string {
	result := []rune(s)

	for i, char := range result {
		if char >= 'a' && char <= 'z' { // harfleri tanımladık.
			result[i] = 'a' + ((char - 'a' + 14) % 26) // 26 Tane harf olan alfabe icin her harfi 14 satır öteledik.
		} else if char >= 'A' && char <= 'Z' {
			result[i] = 'A' + ((char - 'A' + 14) % 26) // büyük harfler icinde öteleme yaptık.
		}
	}

	return string(result)
}
