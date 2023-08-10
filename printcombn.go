package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	nums := make([]int, n)
	printCombination(nums, n, 0, 0)
}

func printCombination(nums []int, n, index, start int) {
	if index == n {
		printNumber(nums)
		z01.PrintRune('\n')
		return
	}

	for i := start; i <= 9; i++ {
		nums[index] = i
		printCombination(nums, n, index+1, i+1)
	}
}

func printNumber(nums []int) {
	for i, num := range nums {
		if i > 0 {
			if i != 1 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			} else {
				z01.PrintRune('0')
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		printDigit(num)
	}
}

func printDigit(digit int) {
	z01.PrintRune(rune(digit) + '0')
}
