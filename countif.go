package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, value := range tab {
		if f(value) {
			count++
		}
	}
	return count
}
