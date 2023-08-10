package piscine

func StrLen(s string) int {
	count := 0

	for range s {
		count = count + 1
	}
	return count
}
