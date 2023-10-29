package piscine

func Join(s []string, sep string) string {
	concat := ""
	for i, res := range s {
		concat += res
		if i != len(s)-1 {
			concat += sep
		}
	}
	return concat
}
