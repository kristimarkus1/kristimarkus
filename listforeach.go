package piscine

func ListForEach(l *List, f func(*NodeL)) {
	current := l.Head
	for current != nil {
		f(current)
		current = current.Next
	}
}

func Add2_node(node *NodeL) {
	switch data := node.Data.(type) {
	case int:
		node.Data = data + 2
	case string:
		node.Data = data + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch data := node.Data.(type) {
	case int:
		node.Data = data - 3
	case string:
		node.Data = data + "-3"
	}
}
