package piscine

func CharIndex(char byte, base string) int {
	for i := 0; i < len(base); i++ {
		if base[i] == char {
			return i
		}
	}
	return -1
}

func IsValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	charMap := make(map[byte]bool)
	for i := 0; i < len(base); i++ {
		if charMap[base[i]] || base[i] == '+' || base[i] == '-' {
			return false
		}
		charMap[base[i]] = true
	}
	return true
}

func AtoiBase(s string, base string) int {
	if !IsValidBase(base) {
		return 0
	}
	baseLength := len(base)
	result := 0
	for i := 0; i < len(s); i++ {
		digitIndex := CharIndex(s[i], base)
		result = result*baseLength + digitIndex
	}
	return result
}
