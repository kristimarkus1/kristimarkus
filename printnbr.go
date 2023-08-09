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

	reverse := 0
	temp := n

	for temp > 0 {
		remainder := temp % 10
		reverse = reverse*10 + remainder
		temp /= 10
	}

	for reverse > 0 {
		remainder := reverse % 10
		z01.PrintRune(rune(remainder + '0'))
		reverse /= 10
	}
}
