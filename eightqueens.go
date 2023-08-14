// Write a function that prints the solutions to the eight queens puzzle.

// Recursivity must be used to solve this problem.

// It should print something like this :
//$ go run .
//15863724
//16837425
//17468253
//...

package piscine

import (
	"github.com/01-edu/z01"
)

func EightQueens() {
	board := make([]int, 8)
	solveEightQueens(board, 0)
}

func solveEightQueens(board []int, col int) {
	if col >= 8 {
		printBoard(board)
		return
	}

	for row := 0; row < 8; row++ {
		if isSafe(board, row, col) {
			board[col] = row
			solveEightQueens(board, col+1)
			board[col] = 0
		}
	}
}

func isSafe(board []int, row, col int) bool {
	for i := 0; i < col; i++ {
		if board[i] == row || abs(board[i]-row) == abs(i-col) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printBoard(board []int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i] == j {
				z01.PrintRune(rune(j + 49))
			}
		}
	}
	z01.PrintRune('\n')
}
