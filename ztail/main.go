package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . -c NUM FILE1 [FILE2...]")
		os.Exit(1)
	}

	num := 0
	for _, c := range os.Args[2] {
		if c < '0' || c > '9' {
			fmt.Println("Invalid number:", os.Args[2])
			os.Exit(1)
		}
		num = num*10 + int(c-'0')
	}

	exitStatus := 0
	for i, file := range os.Args[3:] {
		if i > 0 {
			fmt.Println()
		}
		if len(os.Args) > 4 {
			fmt.Printf("==> %s <==\n", file)
		}
		f, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			exitStatus = 1
			continue
		}
		defer f.Close()
		fi, err := f.Stat()
		if err != nil {
			fmt.Println(err)
			exitStatus = 1
			continue
		}
		size := fi.Size()
		start := size - int64(num)
		if start < 0 {
			start = 0
		}
		buf := make([]byte, num)
		n, err := f.ReadAt(buf, start)
		if err != nil && err.Error() != "EOF" {
			fmt.Println(err)
			exitStatus = 1
			continue
		}
		fmt.Print(string(buf[:n]))
	}
	os.Exit(exitStatus)
}
