package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for a := '0'; a <= '7'; a++ {
		for b := a + 1; b <= '8'; b++ {
			for c := b + 1; c <= '9'; c++ {
				if a != b && b != c {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)

					if a != '7' || b != '8' || c != '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintComb()
}
