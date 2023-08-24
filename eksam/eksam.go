//-------------02-------------//
//Write a program that prints the Latin alphabet in lowercase on a single line.
//A line is a sequence of characters preceding the end of line character ('\n').
package main

import "github.com/01-edu/z01"

func main(){
	for i := 'a' ; i <= 'z' ; i ++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
//Write a program that prints the Latin alphabet in lowercase in reverse order (from 'z' to 'a') on a single line.
//A line is a sequence of characters preceding the end of line character ('\n').
//Please note that casting is not allowed for this exercise!
package main

import "github.com/01-edu/z01"

func main(){
	for i := 'z' ; i <= 'a' ; i ++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
//Write a program that prints the decimal digits in ascending order (from 0 to 9) on a single line.
//A line is a sequence of characters preceding the end of line character ('\n')
package main

import "github.com/01-edu/z01"

func main(){
	for i := '0' ; i <= '9' ; i ++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
//Write a function that prints 'T' (true) 
//on a single line if the int passed as parameter is negative, otherwise it prints 'F' (false).
package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int){
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}
//Write a function that prints, in ascending order and on a single line: all unique combinations of 
//three different digits so that, the first digit is lower than the second, and the second is lower than the third.
//These combinations are separated by a comma and a space.
package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 9; j++ {
			for k := j + 1; k < 10; k++ {
				z01.PrintRune(rune(i + '0'))
				z01.PrintRune(rune(j + '0'))
				z01.PrintRune(rune(k + '0'))
				if i < 7 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
//-------------03-------------//
//Write a function that takes a pointer to an int as argument and gives to this int the value of 1.
package piscine

func PointOne(n *int) {
	*n = 1
}
//Write a function that takes a pointer to a pointer to a pointer to an 
//int as argument and gives to this int the value of 1.
package piscine

func PointOne(n ***int) {
	***n = 1
}
//This function will divide the int a and b.
//The result of this division will be stored in the int pointed by div.
//The remainder of this division will be stored in the int pointed by mod
package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b	 
}
//UltimateDivMod should divide the dereferenced value of a by the dereferenced value of b.
//Store the result of the division in the int which a points to.
//Store the remainder of the division in the int which b points to.
package piscine

func UltimateDivMod(a *int, b *int) {
	*a, *b = *a / *b, *a % *b
}
//Write a function that prints one by one the characters of a string on the screen.
package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	sString := []byte(s)

	for _, abc := range sString {
		z01.PrintRune(rune(abc))
	}
}
//Write a function that counts the runes of a string and that returns that count
package piscine

func StrLen(s string) int {
	count := 0

	for range s {
		count = count + 1
	}
	return count
}
//Write a function that takes two pointers to an int (*int) and swaps their contents
package piscine

func Swap(a *int, b *int) {
	kohatäitja := *a
	*a = *b
	*b = *kohatäitja
}
//Write a function that reverses a string.
//This function will return the reversed string
package piscine

func StrRev(s string) string {
	var result string
	sString := []byte(s)

	for _, abc := range sString {
		result = string(abc) + result
	}
	return result
}
//-------------04-------------//
//Write an iterative function that returns the factorial of the int passed as parameter.
//Errors (non possible values or overflows) will return 0.
package piscine
func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 21 {
		return 0
	}
	for i := 1; i < nb+1; i++ {
		result = result * i
	}
	return result
}
//Write a recursive function that returns the factorial of the int passed as parameter.
//Errors (non possible values or overflows) will return 0.
//for is forbidden for this exercise.
package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}

	if nb == 1 || nb == 0 {
		return 1
	}

	if nb > 1 {
		return nb * RecursiveFactorial(nb-1)
	}
	return 0
}
//Write an iterative function that returns the value of nb to the power of power.
//Negative powers will return 0. Overflows do not have to be dealt with.
package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	result := 1
	for i := 0; i < power; i++ {
		result = result * nb
	}
	return result
}
//Write a recursive function that returns the value of nb to the power of power.
//Negative powers will return 0. Overflows do not have to be dealt with.
//for is forbidden for this exercise.
package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}

	if nb > 0 {
		return nb * RecursivePower(nb, power-1)
	}
	return nb * RecursivePower(nb, power-1)
}
//Write a recursive function that returns the value at the position index in the fibonacci sequence.
//The first value is at index 0.
//The sequence starts this way: 0, 1, 1, 2, 3 etc...
//A negative index will return -1.
//for is forbidden for this exercise.
package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}

	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}
//-------------05-------------//
//Write a function that returns the first rune of a string
package piscine

func FirstRune(s string) rune {
	for _, letter := range s {
		return rune(letter)
	}
	return 0
}
//Write a function that returns the last rune of a string
package piscine

func LastRune(s string) rune {
	lastRune := rune(0)

	for _, letter := range s {
		lastRune = rune(letter)
	}

	return lastRune
}
//Write a function that returns the nth rune of a string. If not possible, it returns 0.
package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)

	if n <= 0 || n > len(runes) {
		return 0
	}
	return runes[n-1]
}
//Write a function that behaves like the Compare function
package piscine

func Compare(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
//Write a function that counts the letters of a string and returns the count.
//The letters are from the Latin alphabet list only, any other characters, symbols or empty spaces 
//shall not be counted.
package piscine

func isLatinLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func AlphaCount(s string) int {
	counter := 0

	for _, letter := range s {
		if isLatinLetter(letter) {
			counter++
		}
	}
	return counter
}
//Write a function that behaves like the Index function.
package piscine

func Index(s string, toFind string) int {
	length := len(toFind)
	length2 := len(s)

	for i := 0; i <= length2-length; i++ {
		if s[i:i+length] == toFind {
			return i
		}
	}
	return -1
}
//Write a function that returns the concatenation of two string passed in arguments.
package piscine

func Concat(str1 string, str2 string) string {
	return str1 + str2
}
//Write a function that returns true if the string passed as 
//parameter contains only uppercase characters, otherwise, returns false
package piscine

func IsUpper(s string) bool {
	for _, letter := range s {
		if letter < 'A' || letter > 'Z' {
			return false
		}
	}
	return true
}
//Write a function that returns true if the string passed as the 
//parameter contains only lowercase characters, otherwise, returns false
package piscine

func IsLower(s string) bool {
	for _, letter := range s {
		if letter < 'a' || letter > 'z' {
			return false
		}
	}
	return true
}
//Write a function that returns true if the string passed as the 
//parameter only contains alphanumerical characters or is empty, otherwise, and returns false
package piscine

func IsAlpha(s string) bool {
	for _, letter := range s {
		if !(letter >= 'a' && letter <= 'z') && !(letter >= 'A' && letter <= 'Z') && !(letter >= '0' && letter <= '9') {
			return false
		}
	}
	return true
}
//Write a function that returns true if the string passed as a 
//parameter contains only numerical characters, otherwise, returns false
package piscine

func IsNumeric(s string) bool {
	for _, letter := range s {
		if !(letter >= '0' && letter <= '9') {
			return false
		}
	}
	return true
}
//Write a function that returns true if the string passed as a 
//parameter contains only printable characters, otherwise, returns false
package piscine

func IsPrintable(s string) bool {
	for _, char := range s {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}
//Write a function that capitalizes each letter in a string
package piscine

func ToUpper(s string) string {
	result := []rune(s)

	for i := range result {
		if result[i] >= 'a' && result[i] <= 'z' {
			result[i] = result[i] - 'a' + 'A'
		}
	}
	return string(result)
}
//Write a function that lower cases for each letter in a string
package piscine

func ToLower(s string) string {
	result := []rune(s)

	for i := range result {
		if result[i] >= 'A' && result[i] <= 'Z' {
			result[i] = result[i] - 'A' + 'a'
		}
	}
	return string(result)
}
//Write a function which prints the digits of an int passed in parameter in ascending order. 
//All possible values of type int have to go through, excluding negative numbers. Conversion to int64 is not allowed
package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	digits := make([]int, 0)
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}

	if len(digits) == 0 {
		z01.PrintRune('0')
		return
	}

	for i := 0; i < len(digits)-1; i++ {
		for j := 0; j < len(digits)-i-1; j++ {
			if digits[j] > digits[j+1] {
				digits[j], digits[j+1] = digits[j+1], digits[j]
			}
		}
	}

	for _, digit := range digits {
		z01.PrintRune(rune(digit) + '0')
	}
}
//Write a function that transforms numbers within a string, into an int.
//If the - sign is encountered before any number it should determine the sign of the returned int.
//This function should only return an int. In the case of an invalid input, the function should return 0.
//Note: There will never be more than one sign in a string in the tests.
package piscine

func TrimAtoi(s string) int {
	var neg bool = false
	var empty bool = true
	var res int = 0

	for _, v := range s {
		if empty && !neg && v == '-' {
			neg = true
		} else if v >= '0' && v <= '9' {
			res *= 10
			res += int(v - '0')
			empty = false
		}
	}
	if empty {
		return 0
	} else {
		if neg {
			return -res
		} else {
			return res
		}
	}
}
//Write a function that capitalizes the first letter of each word and lowercases the rest.
//A word is a sequence of alphanumeric characters
package piscine

func IsAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func Capitalize(s string) string {
	result := []rune(s)
	previousIsAlphaNumeric := false

	for i, char := range result {
		if IsAlphaNumeric(char) {
			if !previousIsAlphaNumeric {
				if char >= 'a' && char <= 'z' {
					result[i] = char - 'a' + 'A'
				}
			} else {
				if char >= 'A' && char <= 'Z' {
					result[i] = char - 'A' + 'a'
				}
			}
			previousIsAlphaNumeric = true
		} else {
			previousIsAlphaNumeric = false
		}
	}

	return string(result)
}
//-------------06-------------//
//Write a program that prints the name of the program.
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]

	for _, character := range arguments[2:] {
		z01.PrintRune(rune(character))
	}
	z01.PrintRune('\n')
}
//Write a program that prints the arguments received in the command line
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
		z01.PrintRune('\n')
	}
}
//Write a program that takes string as arguments, and displays its first argument.
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
//Write a program that prints the arguments received in the command line in reverse order.
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
