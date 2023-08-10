package piscine

func StrLen(s string) int {
	sString := []byte(s)

	count := 0

	for range sString {
		count = count + 1
	}
	return count
}
