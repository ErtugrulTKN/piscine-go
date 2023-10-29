package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{}
	}

	slice := []int{} // Azalan aralığı saklamak için boş bir tamsayı dilimi ekledik.
	for i := max; i > min; i-- {
		slice = append(slice, i)
	}

	return slice
}
