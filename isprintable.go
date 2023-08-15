// Write a function that returns true if the string passed as a parameter
// contains only printable characters, otherwise, returns false.

package piscine

func IsPrintable(s string) bool {
	for _, char := range s {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}
