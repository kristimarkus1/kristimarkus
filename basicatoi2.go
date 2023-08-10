// Write a function that simulates the behaviour of the Atoi function in Go. Atoi transforms a number defined as a string in a number defined as an int.

// Atoi returns 0 if the string is not considered as a valid number. For this exercise non-valid string chains will be tested. Some will contain non-digits characters.

// For this exercise the handling of the signs + or - does not have to be taken into account.

// This function will only have to return the int. For this exercise the error return of Atoi is not required.

package piscine

func BasicAtoi2(s string) int {
	result := 0

	for _, c := range s {
		digit := int(c - '0')
		if digit < 0 || digit > 9 {
			// Skip non-digit characters
			continue
		}
		result = result*10 + digit
	}

	return result
}
