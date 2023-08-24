package piscine

func ListReverse(l *List) {
	var prev *NodeL = nil
	current := l.Head
	l.Tail = l.Head

	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}

	l.Head = prev
}
