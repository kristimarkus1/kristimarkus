// Write a function that capitalizes the first letter
// of each word and lowercases the rest.

package piscine

func IsAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func Capitalize(s string) string {
	result := []rune(s)
	previousIsAlphaNumeric := false

	for i, char := range result {
		if IsAlphaNumeric(char) {
			if !previousIsAlphaNumeric {
				if char >= 'a' && char <= 'z' {
					result[i] = char - 'a' + 'A'
				}
			} else {
				if char >= 'A' && char <= 'Z' {
					result[i] = char - 'A' + 'a'
				}
			}
			previousIsAlphaNumeric = true
		} else {
			previousIsAlphaNumeric = false
		}
	}

	return string(result)
}
