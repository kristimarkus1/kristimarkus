// Write a function that capitalizes the first letter
// of each word and lowercases the rest.

package piscine

package piscine

func Capitalize(s string) string {
	result := ""
	capitalizeNext := true

	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			if capitalizeNext {
				if char >= 'a' && char <= 'z' {
					result += string(char - 'a' + 'A')
				} else {
					result += string(char)
				}
				capitalizeNext = false
			} else {
				result += string(char)
			}
		} else {
			result += string(char)
			capitalizeNext = true
		}
	}

	return result
}

