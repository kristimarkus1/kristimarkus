// Write a program that displays, on the standard output, the content of a file given as argument.
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
	} else {
		fileName := args[0]

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()

		buf := make([]byte, 1024)
		for {
			n, err := file.Read(buf)
			if n > 0 {
				os.Stdout.Write(buf[:n])
			}
			if err != nil {
				break
			}
		}
	}
}
