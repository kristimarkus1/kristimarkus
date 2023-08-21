package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		reader := io.Reader(os.Stdin)
		buffer := make([]byte, 4096)
		for {
			n, err := reader.Read(buffer)
			if n > 0 {
				for _, r := range buffer[:n] {
					z01.PrintRune(rune(r))
				}
			}
			if err == io.EOF {
				break
			}
		}
	} else {
		for _, arg := range args {
			file, err := os.Open(arg)
			if err != nil {
				errMessage := "ERROR: " + err.Error() + "\n"
				for _, r := range errMessage {
					z01.PrintRune(r)
				}
			} else {
				buffer := make([]byte, 4096)
				for {
					n, err := file.Read(buffer)
					if n > 0 {
						for _, r := range buffer[:n] {
							z01.PrintRune(rune(r))
						}
					}
					if err == io.EOF {
						break
					}
				}
				file.Close()
			}
		}
	}
}
