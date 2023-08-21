package main

import (
	"fmt"
	"io"
	"os"
)

func printTail(filename string, count int, multipleFiles bool) {
	if multipleFiles {
		fmt.Printf("==> %s <==\n", filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // Exit with status 1 on error
	}
	defer file.Close()

	buffer := make([]byte, count)
	file.Seek(-int64(count), io.SeekEnd)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		os.Exit(1) // Exit with status 1 on error
	}

	fmt.Print(string(buffer[:n]))

	if multipleFiles {
		fmt.Println()
	}
}

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Usage: go run . -c <count> <file1> [<file2> ...]")
		os.Exit(1)
	}

	option := args[0]
	if option != "-c" {
		fmt.Println("Usage: go run . -c <count> <file1> [<file2> ...]")
		os.Exit(1)
	}

	count := 0
	fmt.Sscanf(args[1], "%d", &count)

	files := args[2:]

	for _, filename := range files {
		printTail(filename, count, len(files) > 1)
	}

	os.Exit(0) // Exit with status 0 on success
}
