package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}

	if l == nil || data_ref <= l.Data {
		newNode.Next = l
		return newNode
	}

	current := l
	for current.Next != nil && data_ref > current.Next.Data {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
	return l
}
