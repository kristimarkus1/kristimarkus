package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	n := len(a)
	ascending := true
	descending := true

	for i := 1; i < n; i++ {
		result := f(a[i-1], a[i])
		if ascending && result > 0 {
			ascending = false
		}
		if descending && result < 0 {
			descending = false
		}
	}

	return ascending || descending
}
