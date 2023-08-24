package piscine

func ListSize(l *List) int {
	size := 0
	current := l.Head
	for current != nil {
		size++
		current = current.Next
	}
	return size
}
