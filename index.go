// Write a function that behaves like the Index function.

package piscine

func Index(s string, toFind string) int {
	length := len(toFind)
	length2 := len(s)

	for i := 0; i <= length2-length; i++ {
		if s[i:i+length] == toFind {
			return i
		}
	}
	return -1
}
