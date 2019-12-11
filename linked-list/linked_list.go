package linkedlist

// Node ...
type Node struct {
	Val          interface{}
	NextNode     *Node
	PreviousNode *Node
}

// List ...
type List struct {
	head *Node
	tail *Node
}

// ErrEmptyList ...
var ErrEmptyList error

// Next ...
func (e *Node) Next() *Node {
	var n *Node
	if e != nil {
		return e.NextNode
	}
	return n
}

// Prev ...
func (e *Node) Prev() *Node {
	var n *Node
	if e != nil {
		return e.PreviousNode
	}

	return n
}

// NewList ...
func NewList(args ...interface{}) *List {
	newList := &List{}

	for _, arg := range args {
		newList.PushBack(arg.(int))
	}

	return newList
}

// PushFront ...
func (l *List) PushFront(v interface{}) {

	if l.tail == nil {
		l.tail = &Node{
			PreviousNode: nil,
			Val:          v.(int),
			NextNode:     nil,
		}
		l.head = l.tail
	} else {
		newNode := &Node{
			PreviousNode: nil,
			Val:          v.(int),
			NextNode:     l.head,
		}
		l.head.PreviousNode = newNode
		l.head = newNode
	}

}

// PushBack ...
func (l *List) PushBack(v interface{}) {
	if l.head == nil {
		l.head = &Node{
			PreviousNode: nil,
			Val:          v.(int),
			NextNode:     nil,
		}

		l.tail = l.head
	} else {
		newNode := &Node{
			PreviousNode: l.tail,
			Val:          v.(int),
			NextNode:     nil,
		}
		l.tail.NextNode = newNode
		l.tail = newNode
	}

}

// PopFront ...
func (l *List) PopFront() (interface{}, error) {

	if l.head == nil {
		return nil, ErrEmptyList
	}

	firstNode := l.head
	var secNode *Node
	secNode = nil

	if l.head.NextNode != nil {
		secNode = l.head.Next()
		secNode.PreviousNode = nil
	}

	l.head = secNode

	if l.head == nil {
		l.tail = l.head
	}

	return firstNode.Val, nil
}

// PopBack ...
func (l *List) PopBack() (interface{}, error) {

	if l.tail == nil {
		return nil, ErrEmptyList
	}

	firstNode := l.tail
	var secNode *Node
	secNode = nil

	if l.tail.PreviousNode != nil {
		secNode = l.tail.PreviousNode
		secNode.NextNode = nil
	}

	l.tail = secNode

	if l.tail == nil {
		l.head = l.tail
	}

	return firstNode.Val, nil
}

// Reverse ...
func (l *List) Reverse() *List {

	if l.head == nil || l.head == l.tail {
		return l
	}

	node := l.tail

	for node.PreviousNode != nil {
		tmp := node.PreviousNode
		node.PreviousNode = node.NextNode
		node.NextNode = tmp

		node = tmp
	}

	tmp := node.PreviousNode
	node.PreviousNode = node.NextNode
	node.NextNode = tmp

	tmp = l.tail
	l.tail = l.head
	l.head = tmp

	return l
}

// First ..
func (l *List) First() *Node {
	return l.head
}

// Last ...
func (l *List) Last() *Node {
	return l.tail
}
