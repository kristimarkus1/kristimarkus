package main

import (
	"fmt"
	"os"
)

func ConvertToInteger(option []rune) int {
	value := 0
	for i := 0; i < len(option); i++ {
		digit := int(option[i] - '0')
		decimal := 1
		for j := 0; j < len(option)-1-i; j++ {
			decimal *= 10
		}
		value += digit * decimal
	}
	return value
}

func main() {
	status := true
	if len(os.Args) >= 4 {
		option := []rune(os.Args[2])
		value := ConvertToInteger(option)
		for i := 3; i < len(os.Args); i++ {
			file, err := os.Open(os.Args[i])
			if err != nil {
				fmt.Printf("open %v: no such file or directory\n", os.Args[i])
				status = false
				continue
			}
			if len(os.Args) > 4 {
				if i > 3 {
					fmt.Print("\n")
				}
				fmt.Printf("==> %v <==\n", os.Args[i])
			}
			file_stat, err := file.Stat()
			if err != nil {
				fmt.Printf("could not obtain stat for file %v\n", os.Args[i])
			}
			content := make([]byte, file_stat.Size())
			file.Read(content)
			contentinrune := []rune(string(content))
			if len(contentinrune) >= value {
				last_chars := make([]rune, value)
				for j := 0; j < len(last_chars); j++ {
					last_chars[j] = contentinrune[len(contentinrune)-value+j]
				}
				fmt.Print(string(last_chars))
			} else {
				fmt.Print(string(contentinrune))
			}
		}
	} else {
		status = false
	}
	if !status {
		os.Exit(1)
	}
}
