package main

import (
	"fmt"
	"os"
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

	success := true

	for i, filename := range files {
		if i > 0 {
			fmt.Println()
		}
		if !printTail(filename, count) {
			success = false
		}
	}

	if !success {
		os.Exit(1)
	}
}

func usage() {
	// Print usage message and exit
	fmt.Printf("Usage: %s -c <count> <file1> [<file2> ...]\n", os.Args[0])
	os.Exit(1)
}

func printTail(filename string, count int) bool {
	file, err := os.Open(filename)
	if err != nil {
		errMessage := fmt.Sprintf("open %s: %s\n", filename, err.Error())
		fmt.Print(errMessage)
		return false
	}
	defer file.Close()

	stat, _ := file.Stat()
	if stat.Size() < int64(count) {
		errMessage := fmt.Sprintf("%s\n", "File size is smaller than count")
		fmt.Print(errMessage)
		return false
	}
	file.Seek(stat.Size()-int64(count), 0) // Seek to the beginning of the count bytes from the end
	buffer := make([]byte, count)
	n, err := file.Read(buffer)
	if err != nil {
		errMessage := fmt.Sprintf("%s\n", err.Error())
		fmt.Print(errMessage)
		return false
	}

	fmt.Printf("==> %s <==\n%s\n", filename, buffer[:n])

	return true
}
