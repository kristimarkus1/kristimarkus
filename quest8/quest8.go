package main

import (
	"github.com/01-edu/z01"
)

type student struct {
	name string
	age int
}

func main() {

	chris  :=  student{}

	z01.PrintRune(rune(chris))
}