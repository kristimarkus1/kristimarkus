package main

import "github.com/01-edu/z01"

type numb struct {
	x int
	y int
}

func setPoint(ptr *numb, x int, y int) {
	ptr.x = x
	ptr.y = y
}

func main() {
	var points numb

	setPoint(&points, 42, 21)

	for _, r := range "x = " {
		z01.PrintRune(r)
	}

	// z01.PrintRune(rune(points.x + 42))
	z01.PrintRune(rune('0' + points.x/10))
	z01.PrintRune(rune('0' + points.x%10))

	for _, i := range ", y = " {
		z01.PrintRune(i)
	}

	// z01.PrintRune(rune(points.y) + 21)
	z01.PrintRune(rune('0' + points.y/10))
	z01.PrintRune(rune('0' + points.y%10))

	z01.PrintRune('\n')
}

/*$ go run .
x = 42, y = 21
$*/
