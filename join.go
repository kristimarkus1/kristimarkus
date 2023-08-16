package piscine

func Join(strs []string, sep string) string {
	result := ""
	for i, s := range strs {
		result += s
		if i < len(strs)-1 {
			result += sep
		}
	}
	return result
}
