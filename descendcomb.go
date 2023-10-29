package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 99; i >= 1; i-- { // ilk sayının iki basamağı
		for j := i - 1; j >= 0; j-- { // ikinci sayının iki basamağı
			PrintNumbers(i/10, i%10)
			PrintNumbers()
			PrintNumbers(j/10, j%10)

			if i != 1 || j != 0 { // Son iki sayı 01 00 olmalı
				z01.PrintRune(',')
				PrintNumbers()
			}
		}
	}
}

func PrintNumbers(numbers ...int) {
	if len(numbers) == 0 { // sıfıra eşitse boş
		z01.PrintRune(' ')
		return
	}

	for _, number := range numbers {
		z01.PrintRune(rune('0' + number))
	}
}
