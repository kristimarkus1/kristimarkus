package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) != 4 {
		return
	}

	arg1, err := strconv.Atoi(args[1])
	if err != nil {
		return
	}

	arg2, err := strconv.Atoi(args[3])
	if err != nil {
		return
	}

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
			fmt.Println("No division by 0")
			return
		}
	} else if operator == "%" {
		if arg2 != 0 {
			result = arg1 % arg2
		} else {
			fmt.Println("No modulo by 0")
			return
		}
	} else {
		validOperator = false
	}

	if validOperator {
		fmt.Println(result)
	}
}
