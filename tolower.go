// Write a function that lower cases for each letter in a string.

package piscine

func ToLower(s string) string {
	result := []rune(s)

	for i := range result {
		if result[i] >= 'A' && result[i] <= 'Z' {
			result[i] = result[i] - 'A' + 'a'
		}
	}
	return string(result)
}
