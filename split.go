package piscine

func Split(s, text string) []string {
	lenOftext := len(text)
	lenOfs := len(s)

	for i := 0; i < lenOfs-lenOftext; i++ {
		if s[i:i+lenOftext] == text {
			s = s[:i] + " " + s[i+lenOftext:]
			lenOfs -= lenOftext
		}
	}

	return SplitWhiteSpaces(s)
}
