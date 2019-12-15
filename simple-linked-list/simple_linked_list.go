package linkedlist

import (
	"errors"
)

// API:
//
type Element struct {
	data int
	next *Element
}

// List ...
type List struct {
	head *Element
	size int
}

// New ...
func New(arr []int) *List {

	l := &List{nil, 0}

	if len(arr) > 0 {
		for _, i := range arr {
			l.Push(i)
		}
	}

	return l
}

// Size ...
func (l *List) Size() int {

	if l == nil {
		return 0
	}

	return l.size
}

// Push ...
func (l *List) Push(i int) {

	newElement := &Element{
		data: i,
		next: nil,
	}

	if l.head == nil {
		l.head = newElement
		l.size++
	} else {
		element := l.head
		for element.next != nil {
			element = element.next
		}
		element.next = newElement
		l.size++
	}
}

// Pop ...
func (l *List) Pop() (int, error) {
	var data int

	if l.head == nil {
		return 0, errors.New("")
	} else if l.head.next == nil {
		data = l.head.data
		l.head = nil
		l.size--
		return data, nil
	}

	element := l.head

	for element.next.next != nil {
		element = element.next
	}

	el := element.next

	element.next = nil
	l.size--
	return el.data, nil

}

// Array ...
func (l *List) Array() []int {
	var arr []int

	element := l.head

	for element != nil {
		arr = append(arr, element.data)
		element = element.next
	}

	return arr
}

// Reverse ..
func (l *List) Reverse() *List {

	if l == nil || l.head == nil {
		return l
	}

	previous := l.head
	curr := l.head.next

	l.head.next = nil

	for curr != nil {
		tmp := curr.next

		curr.next = previous

		previous = curr
		l.head = curr
		curr = tmp
	}

	return l

}
