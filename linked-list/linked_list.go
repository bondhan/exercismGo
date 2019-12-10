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
var ErrEmptyList error

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
		newList.PushFront(arg.(int))
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
			PreviousNode: l.tail,
			Val:          v.(int),
			NextNode:     nil,
		}
		l.tail.NextNode = newNode
		l.tail = newNode
	}

}

// PushBack ...
func (l *List) PushBack(v interface{}) {
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

	if l.head == nil || l.head == l.tail {
		return l
	}

	firstNode := l.tail
	secondNode := firstNode.PreviousNode

	for firstNode.PreviousNode != nil {
		firstNode.NextNode = secondNode
		secondNode.PreviousNode = firstNode
	}

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
