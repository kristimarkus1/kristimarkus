// Write a function that returns true if the string passed as the parameter only
// contains alphanumerical characters or is empty, otherwise, and returns false.

package piscine

func IsAlpha(s string) bool {
	for _, letter := range s {
		if !(letter >= 'a' && letter <= 'z') && !(letter >= 'A' && letter <= 'Z') && !(letter >= '0' && letter <= '9') {
			return false
		}
	}
	return true
}
