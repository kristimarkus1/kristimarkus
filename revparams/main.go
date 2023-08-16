// Write a program that prints the arguments received in the command line in reverse order.

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	length := 0

	for index := range arguments {
		length = index
	}

	for i := length; i > 0; i-- {
		for _, all := range arguments[i] {
			z01.PrintRune(all)
		}
		z01.PrintRune('\n')
	}
}
