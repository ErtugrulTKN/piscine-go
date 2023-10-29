package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}

	if len(toFind) > len(s) {
		return -1
	}

	for i := 0; i <= len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
