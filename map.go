package piscine

func Map(f func(int) bool, a []int) []bool { // Bir int dilimi için, bu dilimin her bir öğesine bool türünde bir işlev uygulayan fonksiyon.

	result := make([]bool, len(a)) // resault bool olmalı
	for i, val := range a {
		result[i] = f(val)
	}
	return result
}
