package piscine

func Compare(a, b string) int {
	A, B := len(a), len(b)
	minLen := A
	if B < minLen {
		minLen = B
	}

	for i := 0; i < minLen; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}

	if A < B {
		return -1
	} else if A > B {
		return 1
	}

	return 0
}
