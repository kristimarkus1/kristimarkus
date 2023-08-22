package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, value := range a {
		result[i] = f(value)
	}
	return result
}
