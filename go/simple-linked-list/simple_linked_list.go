package linkedlist

import (
	"errors"
)

// Define the List and Element types here.
type Node struct {
	next *Node
	data int
}

type List struct {
	head *Node
	size int
}

func New(elements []int) *List {
	l := &List{}
	for _, elm := range elements {
		l.Push(elm)
	}
	return l
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	node := &Node{data: element}
	if l.head == nil {
		l.head = node
		l.size++
	} else {
		curr := l.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = node
		l.size++
	}
}

func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("list is empty")
	}

	if l.head.next == nil {
		data := l.head.data
		l.head = nil
		l.size = 0
		return data, nil
	}
	curr := l.head
	for curr.next.next != nil {
		curr = curr.next
	}
	data := curr.next.data

	curr.next = nil
	l.size--

	return data, nil
}

func (l *List) Array() []int {
	var arr []int
	curr := l.head
	for curr != nil {
		arr = append(arr, curr.data)
		curr = curr.next
	}
	return arr
}

func (l *List) Reverse() *List {
	var prev, next *Node
	current := l.head
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
	return l
}
