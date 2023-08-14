// Write a recursive function that returns the value at the position index in the fibonacci sequence.

// The first value is at index 0.

// The sequence starts this way: 0, 1, 1, 2, 3 etc...

// A negative index will return -1.

// for is forbidden for this exercise.

package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}

	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}
