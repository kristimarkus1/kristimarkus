// Write a function which prints the digits of an int passed in
// parameter in ascending order. All possible values of type int
// have to go through, excluding negative numbers. Conversion to int64 is not allowed.

package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	digits := make([]int, 0)
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}

	if len(digits) == 0 {
		z01.PrintRune('0')
		return
	}

	for i := 0; i < len(digits)-1; i++ {
		for j := 0; j < len(digits)-i-1; j++ {
			if digits[j] > digits[j+1] {
				digits[j], digits[j+1] = digits[j+1], digits[j]
			}
		}
	}

	for _, digit := range digits {
		z01.PrintRune(rune(digit) + '0')
	}
}
