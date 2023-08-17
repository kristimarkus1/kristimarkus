// Write a function that receives three arguments:

// nbr: A string representing a numberic value in a base.

// baseFrom: A string representing the base nbr it's using.

// baseTo: A string representing the base nbr should be represented in the returned value.

// Only valid bases will be tested.

// Negative numbers will not be tested.

package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimalValue := AtoiBase(nbr, baseFrom)
	return NbrBase(decimalValue, baseTo)
}

func NbrBase(n int, base string) string {
	if len(base) < 2 || !IsValidBase(base) {
		return "0"
	}

	result := ""
	baseLength := len(base)

	for n > 0 {
		remainder := n % baseLength
		result = string(base[remainder]) + result
		n /= baseLength
	}

	return result
}
