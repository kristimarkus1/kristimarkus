package piscine

func ListPushFront(l *List, data interface{}) {
	HlnewNode := &NodeL{
		Data: data,
		Next: l.Head,
	}

	l.Head = HlnewNode

	if l.Tail == nil {
		l.Tail = HlnewNode
	}
}
