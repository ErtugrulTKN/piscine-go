package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		// Komut satırı argümanlarından herhangi birinin belirtilen dizelerle eşleşip eşleşmediğini kontrol ettik..
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
