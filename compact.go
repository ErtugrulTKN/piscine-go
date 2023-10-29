package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	newSize := 0

	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			slice[newSize] = slice[i]
			newSize++
		}
	}

	*ptr = slice[:newSize]
	return newSize
}
