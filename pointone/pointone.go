package main

import "fmt"

func PointOne(n *int) {
	*n = 1
}

func main() {
	a := 1

	PointOne(&a)

	var pointer *int = &a

	fmt.Println(*pointer)
}
