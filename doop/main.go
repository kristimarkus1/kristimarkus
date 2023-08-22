package main

import (
	"os"
)

func main() {
	args := os.Args

	if len(args) != 4 {
		return
	}

	arg1 := atoi(args[1])
	arg2 := atoi(args[3])
	operator := args[2]

	result := 0
	validOperator := true

	if operator == "+" {
		result = arg1 + arg2
	} else if operator == "-" {
		result = arg1 - arg2
	} else if operator == "*" {
		result = arg1 * arg2
	} else if operator == "/" {
		if arg2 != 0 {
			result = arg1 / arg2
		} else {
			os.Stdout.Write([]byte("No division by 0\n"))
			return
		}
	} else if operator == "%" {
		if arg2 != 0 {
			result = arg1 % arg2
		} else {
			os.Stdout.Write([]byte("No modulo by 0\n"))
			return
		}
	} else {
		validOperator = false
	}

	if validOperator {
		printNbr(result)
	}
}

func atoi(s string) int {
	neg := false
	num := 0

	for _, c := range s {
		if c == '-' {
			neg = true
		} else if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		} else {
			break
		}
	}

	if neg {
		num *= -1
	}

	return num
}

func printNbr(n int) {
	neg := false

	if n < 0 {
		os.Stdout.Write([]byte("-"))
		neg = true
		n *= -1
	}

	digit := 1
	tmp := n

	for tmp >= 10 {
		tmp /= 10
		digit *= 10
	}

	for digit > 0 {
		digitChar := byte(n/digit + '0')
		os.Stdout.Write([]byte{digitChar})
		n %= digit
		digit /= 10
	}

	if neg {
		n *= -1
	}
}
