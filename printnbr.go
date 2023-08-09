package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == -2147483648 {
		z01.PrintRune('-')
		z01.PrintRune('2')
		printPositiveNbr(147483648)
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	printPositiveNbr(n)
}

func printPositiveNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n > 9 {
		printPositiveNbr(n / 10)
	}

	z01.PrintRune(rune(n%10 + '0'))
}
