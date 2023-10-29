package piscine

func Any(B func(string) bool, A []string) bool {
	for _, s := range A {
		if B(s) {
			return true
		}
	}
	return false
}
