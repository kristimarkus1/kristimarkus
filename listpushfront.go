package piscine

type NodeLo struct {
	Data interface{}
	Next *NodeLo
}

type Listo struct {
	Head *NodeLo
	Tail *NodeLo
}

func ListPushFront(l *Listo, data interface{}) {
	HlnewNode := &NodeLo{
		Data: data,
		Next: l.Head,
	}

	l.Head = HlnewNode

	if l.Tail == nil {
		l.Tail = HlnewNode
	}
}
