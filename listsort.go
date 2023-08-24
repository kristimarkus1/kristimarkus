package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	var sorted *NodeI
	current := l

	for current != nil {
		next := current.Next
		sorted = insertInSortedOrder(sorted, current)
		current = next
	}

	return sorted
}

func insertInSortedOrder(sorted, node *NodeI) *NodeI {
	if sorted == nil || sorted.Data >= node.Data {
		node.Next = sorted
		return node
	}

	current := sorted
	for current.Next != nil && current.Next.Data < node.Data {
		current = current.Next
	}
	node.Next = current.Next
	current.Next = node
	return sorted
}
