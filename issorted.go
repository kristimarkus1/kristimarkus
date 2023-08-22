package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	n := len(a)
	for i := 1; i < n; i++ {
		if f(a[i-1], a[i]) > 0 {
			return false
		}
	}
	return true
}
