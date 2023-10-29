package piscine

func RockAndRoll(n int) string {
	if n < 0 {
		return "error: number is negative\n" // Sayı negatif olmamalı.
	}
	if n%2 == 0 && n%3 == 0 { // Hem 2 hem 3 e bölünebiliyor yani kalan sıfır ise, yazdır.
		return "rock and roll\n"
	} else if n%2 == 0 { // Sadece 2'ye bölünebiliyorsa, yazdır.
		return "rock\n"
	} else if n%3 == 0 { // Sadece 3'e bölünebiliyorsa, yazdır.
		return "roll\n"
	} else {
		return "error: non divisible\n"
	}
}
