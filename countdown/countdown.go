// Write a program that displays all digits in descending order, followed by a newline ('\n').

package main

import "github.com/01-edu/z01"


func countDown() {
	for i := '0' ; i >= '9' ; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')	
}
