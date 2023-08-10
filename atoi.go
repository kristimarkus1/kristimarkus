// Write a function that simulates the behaviour of the Atoi function in Go. Atoi transforms a number represented as a string in a number represented as an int.

// Atoi returns 0 if the string is not considered as a valid number. For this exercise non-valid string chains will be tested. Some will contain non-digits characters.

// For this exercise the handling of the signs + or - does have to be taken into account.

// This function will only have to return the int. For this exercise the error result of Atoi is not required.

package piscine

func Atoi(s string) int {
	result := 0
	sign := 1
	valid := true

	for i, c := range s {
		if i == 0 && (c == '-' || c == '+') {
			if c == '-' {
				sign = -1
			}
			continue
		}

		if c >= '0' && c <= '9' {
			digit := int(c - '0')
			result = result*10 + digit
		} else {
			valid = false
			break
		}
	}

	if !valid {
		return 0
	}

	return result * sign
}
