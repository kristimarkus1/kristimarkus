// Write a function that receives a string slice and prints each element of the slice in a seperate line.

package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, word := range a {
		for _, c := range word {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
