package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	myProgramName := os.Args[0]
	s := []rune(myProgramName)
	for index, name := range s {
		if index > 1 {
			z01.PrintRune(name)
		}
	}
	z01.PrintRune('\n')
}
