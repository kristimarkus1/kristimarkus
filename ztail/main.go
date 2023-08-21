package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 || args[0] != "-c" {
		usage()
	}

	count := 0
	_, err := fmt.Sscanf(args[1], "%d", &count)
	if err != nil {
		usage()
	}

	files := args[2:]

	hasError := false

	for _, filename := range files {
		if len(files) > 1 {
			fmt.Printf("==> %s <==\n", filename)
		}
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("open %s: %s\n", filename, err.Error())
			hasError = true
			continue
		}
		defer file.Close()

		stat, _ := file.Stat()
		file.Seek(stat.Size()-int64(count), 0)
		buffer := make([]byte, count)
		n, err := file.Read(buffer)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			hasError = true
			continue
		}

		fmt.Print(string(buffer[:n]))

		if len(files) > 1 {
			fmt.Print("\n")
		}
	}

	if hasError {
		os.Exit(1)
	}
}

func usage() {
	fmt.Printf("Usage: %s -c <count> <file1> [<file2> ...]\n", os.Args[0])
	os.Exit(1)
}
