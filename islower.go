// Write a function that returns true if the string passed as the parameter
// contains only lowercase characters, otherwise, returns false.

package piscine

func IsLower(s string) bool {
	for _, letter := range s {
		if letter < 'a' || letter > 'z' {
			return false
		}
	}
	return true
}
