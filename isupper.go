// Write a function that returns true if the string passed as parameter
// contains only uppercase characters, otherwise, returns false.

package piscine

func IsUpper(s string) bool {
	for _, letter := range s {
		if letter < 'A' || letter > 'Z' {
			return false
		}
	}
	return true
}
