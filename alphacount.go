// Write a function that counts the letters of a string and returns the count.

// The letters are from the Latin alphabet list only, any other characters, symbols or empty spaces shall not be counted.
package piscine

func isLatinLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func AlphaCount(s string) int {
	counter := 0

	for _, letter := range s {
		if isLatinLetter(letter) {
			counter++
		}
	}
	return counter
}
