package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	var insertString, argumentString string
	var orderFlag bool

	for i, arg := range args {
		switch arg {
		case "--insert", "-i":
			if i+1 < len(args) {
				insertString = args[i+1]
			}
		case "--order", "-o":
			orderFlag = true
		default:
			argumentString += arg + " "
		}
	}

	if insertString != "" {
		argumentString = replaceFirstOccurrence(argumentString, "--insert "+insertString, insertString)
		argumentString = replaceFirstOccurrence(argumentString, "-i "+insertString, insertString)
	}

	if orderFlag {
		argumentString = orderASCII(argumentString)
	}

	printString(argumentString)
}

func printHelp() {
	lines := []string{
		"Options:",
		"  --insert string, -i string",
		"    Insert the string given to the --insert flag into the argument string.",
		"  --order, -o",
		"    Order the string argument in ASCII order.",
	}
	for _, line := range lines {
		printString(line)
	}
}

func orderASCII(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[j] < runes[i] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

func replaceFirstOccurrence(str, old, new string) string {
	if i := indexOf(str, old); i != -1 {
		return str[:i] + new + str[i+len(old):]
	}
	return str
}

func indexOf(str, substr string) int {
	n := len(str)
	m := len(substr)
	for i := 0; i <= n-m; i++ {
		if str[i:i+m] == substr {
			return i
		}
	}
	return -1
}

func printString(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
