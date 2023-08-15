// Write a function that capitalizes each letter in a string.
package piscine

func ToUpper(s string) string {
	result := []rune(s)

	for i := range result {
		if result[i] >= 'a' && result[i] <= 'z' {
			result[i] = result[i] - 'a' + 'A'
		}
	}
	return string(result)
}
