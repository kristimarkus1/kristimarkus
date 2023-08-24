package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	lnewNode := &NodeL{
		Data: data,
		Next: nil,
	}
	if l.Head == nil {
		l.Head = lnewNode
		l.Tail = lnewNode
	} else {
		l.Tail.Next = lnewNode
		l.Tail = lnewNode
	}
}
