// Write a function that separates the words of a string
// and puts them in a string slice.

// The separators are spaces, tabs and newlines.

package piscine

func isWhiteSpace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\n'
}

func SplitWhiteSpaces(s string) []string {
	var result []string
	word := ""
	for _, c := range s {
		if isWhiteSpace(c) {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(c)
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
