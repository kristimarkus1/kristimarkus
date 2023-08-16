// Write a program that prints the name of the program.

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]

	z01.PrintRune([]rune(arguments)[0])
}
