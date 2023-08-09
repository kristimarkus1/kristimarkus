package piscine

import (
	"github.com/01-edu/z01"
)

func printPositiveNbr(n int) {
	if n > 9 {
		printPositiveNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	printPositiveNbr(n)
}
