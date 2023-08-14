package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	result := 1
	for i := 1; i < nb+1; i++ {
		result = result * i
	}
	return result
}
