// Write a program that displays, on the standard output, the content of a file given as argument.
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("File name missing")
		return
	}

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	data := make([]byte, 4096) // Read up to 4096 bytes at a time
	for {
		n, err := file.Read(data)
		if err != nil {
			break
		}
		fmt.Print(string(data[:n]))
	}
}
