package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	sString := []byte(s)

	for _, abc := range sString {
		z01.PrintRune(rune(abc))
	}
}
