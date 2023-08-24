// Write a function that transforms numbers within a string, into an int.

// If the - sign is encountered before any number it should determine the sign of the returned int.

// This function should only return an int. In the case of an invalid input, the function should return 0.

// Note: There will never be more than one sign in a string in the tests.

package piscine

func TrimAtoi(s string) int {
	var neg bool = false
	var empty bool = true
	var res int = 0

	for _, v := range s {
		if empty && !neg && v == '-' {
			neg = true
		} else if v >= '0' && v <= '9' {
			res *= 10
			res += int(v - '0')
			empty = false
		}
	}
	if empty {
		return 0
	} else {
		if neg {
			return -res
		} else {
			return res
		}
	}
}



