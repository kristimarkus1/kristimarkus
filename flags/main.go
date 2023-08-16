package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func PrintInfo() {
	// printing this when there is no args or --help -h
	fmt.Print("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func Sort(s string) {
	var array [1000]int
	for _, char := range s {
		array[int(char)]++
	}
	for index, char := range array {
		for char > 0 {
			z01.PrintRune(rune(index))
			char--
		}
	}
	z01.PrintRune('\n')
}

func main() {
	argArray := os.Args[1:]
	// empty string
	str := ""
	check := false
	SortIt := false
	for _, element := range argArray {
		check = true
		if element == "-h" || element == "--help" {
			PrintInfo() // print info
			break
		}
		// length of an element
		len := 0
		for i := range element {
			len = i + 1
		}
		if len > 0 {
			if element[0] == '-' {
				if len > 2 && element[2] == 'i' {
					// for --insert
					if len > 8 {
						str += element[9:]
					}
				}
				if element[1] == 'i' {
					// for -i
					if len > 3 {
						str += element[3:]
					}
				}
			} else {
				str = element + str
			}
		}
		if element == "-o" || element == "--order" {
			// for --order or -o
			// sending for further "if" block for sorting
			SortIt = true
		}
	}
	// if nothing - print info
	if !check {
		PrintInfo()
	}
	// sorting for --order or -o
	if SortIt {
		Sort(str)
	} else {
		// otherwise print str
		fmt.Println(str)
	}
}
