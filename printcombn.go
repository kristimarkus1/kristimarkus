package piscine

import (
	"github.com/01-edu/z01"
)

// PrintCombN prints all possible combinations of n different digits in ascending order.
func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return // n should be within the range (0, 10)
	}

	nums := make([]int, n)
	printCombination(nums, n, 0, 0)
}

// Helper function to recursively generate and print combinations
func printCombination(nums []int, n, index, start int) {
	if index == n {
		printNumber(nums)
		return
	}

	for i := start; i <= 9; i++ {
		nums[index] = i
		printCombination(nums, n, index+1, i+1)
	}
}

// Helper function to print a combination as a number
func printNumber(nums []int) {
	var buffer []rune

	for i, num := range nums {
		if i > 0 {
			buffer = append(buffer, ',', ' ')
		}
		buffer = append(buffer, rune(num)+'0')
	}
	buffer = append(buffer, ',', ' ')
	buffer = append(buffer, '\n')

	for _, r := range buffer {
		z01.PrintRune(r)
	}
}
