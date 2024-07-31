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
	l.head = &Node{data: element, next: l.head}
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("list is empty")
	}

	deadHead := l.head
	l.head = deadHead.next
	deadHead.next = nil
	l.size--
	return deadHead.data, nil
}

func (l *List) Array() []int {
	output := make([]int, l.size)

	for i, head := l.size-1, l.head; i > -1; i, head = i-1, head.next {
		output[i] = head.data
	}

	return output
}

func (l *List) Reverse() *List {
	output := &List{}
	for head := l.head; head != nil; head = head.next {
		output.Push(head.data)
	}
	return output
}
