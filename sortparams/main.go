package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	// cok anlamadım ama sort algoritması böyleymis
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}

	for _, a := range args {
		for _, b := range a {
			z01.PrintRune(b)
		}
		z01.PrintRune('\n')
	}
}
