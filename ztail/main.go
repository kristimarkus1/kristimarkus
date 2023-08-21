package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 || args[0] != "-c" {
		usage()
	}

	count := 0
	if _, err := fmt.Sscanf(args[1], "%d", &count); err != nil {
		usage()
	}

	files := args[2:]

	for _, filename := range files {
		printTail(filename, count, len(files) > 1)
	}
}

func usage() {
	// Print usage message and exit
	fmt.Println("Usage: go run . -c <count> <file1> [<file2> ...]")
	os.Exit(1)
}

func printTail(filename string, count int, multipleFiles bool) {
	if multipleFiles {
		fmt.Printf("==> %s <==\n", filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		errMessage := "open " + filename + ": " + err.Error() + "\n"
		for _, r := range errMessage {
			z01.PrintRune(r)
		}
		os.Exit(1)
	}
	defer file.Close()

	stat, _ := file.Stat()
	file.Seek(stat.Size()-int64(count), 0) // Seek to the beginning of the count bytes from the end
	buffer := make([]byte, count)
	n, err := file.Read(buffer)
	if err != nil {
		errMessage := "ERROR: " + err.Error() + "\n"
		for _, r := range errMessage {
			z01.PrintRune(r)
		}
		os.Exit(1)
	}

	for _, r := range buffer[:n] {
		z01.PrintRune(rune(r))
	}

	if multipleFiles {
		z01.PrintRune('\n')
	}
}
