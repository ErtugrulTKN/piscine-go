package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	x := 1
	for i := 0; i < 100; i++ {
		x = (x + nb/x) / 2
		if x*x == nb {
			return x
		}
	}
	return 0
}
