package piscine

func SplitWhiteSpaces(texttt string) []string {
	TexteToString := ""
	text := []string{}
	for i, v := range texttt {
		if i == len(texttt)-1 && string(v) != " " && string(v) != "	" && string(v) != "\n" {
			TexteToString += string(v)
			text = append(text, TexteToString)
		} else if string(v) != " " && string(v) != "	" && string(v) != "\n" {
			TexteToString += string(v)
		} else {
			if len(TexteToString) >= 1 {
				text = append(text, TexteToString)
			}
			TexteToString = ""
		}
	}
	return text
}
