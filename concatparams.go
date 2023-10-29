package piscine

func ConcatParams(degisken []string) string {
	str := ""

	for i, res := range degisken {
		str += string(res)
		if i != len(degisken)-1 {
			str += "\n"
		}
	}
	return str
}
