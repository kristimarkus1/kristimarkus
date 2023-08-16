// Write a program that prints the arguments received in the command line.

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	for _, character := range arguments {
		for _, all := range character {
			z01.PrintRune(rune(all))
		}
		z01.PrintRune(' ')
		z01.PrintRune('\n')
	}
}
