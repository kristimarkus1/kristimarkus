package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	var merged *NodeI

	if n1 == nil {
		return n2
	} else if n2 == nil {
		return n1
	}

	if n1.Data <= n2.Data {
		merged = n1
		n1 = n1.Next
	} else {
		merged = n2
		n2 = n2.Next
	}

	current := merged

	for n1 != nil && n2 != nil {
		if n1.Data <= n2.Data {
			current.Next = n1
			n1 = n1.Next
		} else {
			current.Next = n2
			n2 = n2.Next
		}
		current = current.Next
	}

	if n1 != nil {
		current.Next = n1
	} else if n2 != nil {
		current.Next = n2
	}

	return merged
}
