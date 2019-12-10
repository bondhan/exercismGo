package linkedlist

import "github.com/phayes/errors"

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
var ErrEmptyList *List

// Next ...
func (e *Node) Next() *Node {
	return e.NextNode
}

// Prev ...
func (e *Node) Prev() *Node {
	return e.PreviousNode
}

// NewList ...
func NewList(args ...interface{}) *List {
	newList := &List{}

	for _, arg := range args {
		newList.PushFront(arg)
	}

	return newList
}

// PushFront ...
func (l *List) PushFront(v interface{}) {
	if l.head == nil {
		l.head = &Node{
			PreviousNode: nil,
			Val:          v.(int),
			NextNode:     nil,
		}

		l.tail = l.head
	} else {
		newNode := &Node{
			PreviousNode: l.head,
			Val:          v.(int),
			NextNode:     nil,
		}
		l.tail = l.head
		l.head = newNode
	}

}

// PushBack ...
func (l *List) PushBack(v interface{}) {
	if l.tail == nil {
		newNode := &Node{
			PreviousNode: nil,
			Val:          v.(int),
			NextNode:     nil,
		}
		l.tail = newNode
		l.head = newNode
	} else {
		newNode := &Node{
			PreviousNode: nil,
			Val:          v.(int),
			NextNode:     l.tail,
		}
		l.tail = newNode
	}
}

// PopFront ...
func (l *List) PopFront() (interface{}, error) {
	if l.head == nil {
		return ErrEmptyList, errors.New("empty")
	}

	theNode := l.head

	tmpNode := l.head.Prev()
	tmpNode.NextNode = nil
	l.head = tmpNode

	return theNode, nil
}

// PopBack ...
func (l *List) PopBack() (interface{}, error) {
	if l.tail == nil {
		return ErrEmptyList, errors.New("empty")
	}

	theNode := l.tail

	tmpNode := l.tail.Next()
	tmpNode.PreviousNode = nil
	l.tail = tmpNode

	return theNode, nil

}

// Reverse ...
func (l *List) Reverse() *List {
	node := l.head
	for node != nil {
		if node.Prev() != nil {
			node.NextNode = node.Prev()
			node = node.Prev()
		}
	}

	return l
}

// First ..
func (l *List) First() *Node {
	return l.head
}

// Last ...
func (l *List) Last() *Node {
	return l.head
}
