package piscine

func BasicJoin(s []string) string {
	concat := ""
	for _, res := range s {
		concat += res
	}
	return concat
}
