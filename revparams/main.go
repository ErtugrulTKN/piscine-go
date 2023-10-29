package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args
	res := ""
	for _, a := range name[1:] {
		res = a + "\n" + res
	}
	for _, b := range res {
		z01.PrintRune(b)
	}
}
