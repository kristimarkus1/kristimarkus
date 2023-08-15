// Write a function that returns the last rune of a string
package piscine

func LastRune(s string) rune {
	lastRune := rune(0)

	for _, letter := range s {
		lastRune = rune(letter)
	}

	return lastRune
}
