package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args

	for _, word := range arg[1] {
		z01.PrintRune(rune(word))
	}
	z01.PrintRune('\n')
}
