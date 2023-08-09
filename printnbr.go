package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	digits := []int{}
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}

	if len(digits) == 0 {
		z01.PrintRune('0')
	} else {
		for i := len(digits) - 1; i >= 0; i-- {
			z01.PrintRune(rune(digits[i] + '0'))
		}
	}
}
