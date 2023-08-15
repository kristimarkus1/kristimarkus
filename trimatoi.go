// Write a function that transforms numbers within a string, into an int.

// If the - sign is encountered before any number it should determine the sign of the returned int.

// This function should only return an int. In the case of an invalid input, the function should return 0.

// Note: There will never be more than one sign in a string in the tests.

package piscine

func TrimAtoi(s string) int {
	var result int
	sign := 1
	signFound := false

	for _, char := range s {
		if char == '-' && !signFound {
			sign = -1
			signFound = true
		} else if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
			signFound = true
		} else if signFound {
			break
		}
	}

	return result * sign
}
