package piscine

import (
	"fmt"
)

func DealAPackOfCards(deck []int) {
	for i := 0; i < 12; i += 3 { // Her yinelemede 'i'yi 3 artırarak deste üzerinde döngü yapılır.
		player := i/3 + 1 // Oyuncu sayısını hesaplanır Oyuncular 1'den 4'e kadar numaralandırılır
		end := i + 3

		fmt.Printf("Player %d: ", player)
		for i, card := range deck[i:end] { // Mevcut oyuncunun elindeki kartları yinelenir
			fmt.Printf("%d", card)

			if i != 2 {
				fmt.Printf(", ") // Son kart
			}
		}
		fmt.Printf("\n")
	}
}
