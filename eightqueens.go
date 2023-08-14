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
	"fmt"
)

func EightQueens() {
	board := make([][]bool, 8)
	for i := range board {
		board[i] = make([]bool, 8)
	}

	solveEightQueens(board, 0)
}

func solveEightQueens(board [][]bool, col int) {
	if col >= 8 {
		printBoard(board)
		return
	}

	for row := 0; row < 8; row++ {
		if isSafe(board, row, col) {
			board[row][col] = true
			solveEightQueens(board, col+1)
			board[row][col] = false
		}
	}
}

func isSafe(board [][]bool, row, col int) bool {
	// Check the same row to the left
	for i := 0; i < col; i++ {
		if board[row][i] {
			return false
		}
	}

	// Check upper diagonal to the left
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}

	// Check lower diagonal to the left
	for i, j := row, col; i < 8 && j >= 0; i, j = i+1, j-1 {
		if board[i][j] {
			return false
		}
	}

	return true
}

func printBoard(board [][]bool) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] {
				fmt.Printf("%d", j+1)
			}
		}
	}
	fmt.Println()
}
