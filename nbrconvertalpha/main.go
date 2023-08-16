package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		return
	}

	upperCase := false
	if arguments[0] == "--upper" {
		upperCase = true
		arguments = arguments[1:]
	}

	for _, arg := range arguments {
		n := 0
		validNumber := true

		for _, char := range arg {
			if char < '0' || char > '9' {
				validNumber = false
				break
			}
			n = n*10 + int(char-'0')
		}

		if validNumber && n >= 1 && n <= 26 {
			letter := 'a' + rune(n-1)
			if upperCase {
				letter = 'A' + rune(n-1)
			}
			z01.PrintRune(letter)
		} else {
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')
}
