// Write a function that receives a string and a separator and
// returns a slice of strings that results of splitting the string s by the separator sep.

package piscine

func Split(s, sep string) []string {
	var result []string
	start := 0
	for i := 0; i < len(s)-len(sep)+1; i++ {
		if s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			start = i + len(sep)
			i = start - 1 // Adjust i to the next start position
		}
	}
	if start < len(s) {
		result = append(result, s[start:])
	}
	return result
}
