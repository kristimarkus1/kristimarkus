package piscine

func FirstRune(s string) rune {
	for _, letter := range s {
		return rune(letter)
	}
	return 0
}
