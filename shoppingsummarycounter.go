package piscine

func ShoppingSummaryCounter(s string) map[string]int {
	summary := make(map[string]int)
	item := ""

	for A := 0; A <= len(s); A++ {
		if A == len(s) || s[A] == ' ' || s[A] == '\t' || s[A] == '\n' || s[A] == '\r' {
			summary[item]++
			item = ""
			continue
		}

		item += string(s[A])
	}

	return summary
}
