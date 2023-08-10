package piscine

func StrRev(s string) string {
	var result string
	sString := []byte(s)

	for _, abc := range sString {
		result = string(abc) + result
	}
	return result
}
