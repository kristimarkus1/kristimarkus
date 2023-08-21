package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 || args[0] != "-c" {
		usageAndExit()
	}

	count := 0
	_, err := fmt.Sscanf(args[1], "%d", &count)
	if err != nil || count <= 0 {
		usageAndExit()
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

func usageAndExit() {
	os.Exit(1)
}

func printTail(filename string, count int) bool {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open %s: %s\n", filename, err)
		return false
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Printf("%s\n", err)
		return false
	}

	fmt.Printf("==> %s <==\n", filename)

	if stat.Size() < int64(count) {
		fmt.Printf("File size is smaller than count\n")
		return true // Modified this line to return true instead of false
	}

	file.Seek(stat.Size()-int64(count), 0)
	buffer := make([]byte, count)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Printf("%s\n", err)
		return false
	}

	fmt.Printf("%s", buffer)

	return true
}
