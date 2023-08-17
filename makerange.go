// Write a function that takes an int min and an int max as parameters. The function must return a slice of ints with all the values between min and max.

// Min is included, and max is excluded.

// If min is greater than or equal to max, a nil slice is returned.

// append is not allowed for this exercise.
package piscine

func MakeRange(min, max int) []int {
	size := max - min
	if size <= 0 {
		return nil
	}
	answer := make([]int, size)

	for i := 0; i < size; i++ {
		answer[i] = min + i
	}
	return answer
}
