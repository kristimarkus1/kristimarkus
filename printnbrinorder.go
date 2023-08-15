// Write a function which prints the digits of an int passed in
// parameter in ascending order. All possible values of type int
// have to go through, excluding negative numbers. Conversion to int64 is not allowed.

package piscine

import (
	"fmt"
	"sort"
)

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	digits := []int{}
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}

	if len(digits) == 0 {
		fmt.Print("0")
		return
	}

	sort.Ints(digits)
	for _, digit := range digits {
		fmt.Print(digit)
	}
}
