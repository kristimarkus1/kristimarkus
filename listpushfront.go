package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

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
