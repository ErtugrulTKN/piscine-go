package piscine

func ReverseMenuIndex(menu []string) []string { // Bir dizedeki ögeleri tersine çeviren fonksiyon
	out := make([]string, len(menu))

	for i, item := range menu {
		out[len(menu)-i-1] = item
	}

	return out
}
