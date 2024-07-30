package linkedlist

import "errors"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct {
	prev  *Node
	next  *Node
	Value interface{}
}

type List struct {
	head *Node
	tail *Node
}

func NewList(elements ...interface{}) *List {
	l := &List{}
	for _, elem := range elements {
		l.Push(elem)
	}

	return l
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	node := &Node{Value: v}
	headNode := l.head
	if headNode == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = headNode
		headNode.prev = node
		l.head = node
	}
}

func (l *List) Push(v interface{}) {
	node := &Node{Value: v}
	tailNode := l.tail
	if tailNode == nil {
		l.head = node
		l.tail = node
	} else {
		node.prev = tailNode
		tailNode.next = node
		l.tail = node
	}
}

func (l *List) Shift() (interface{}, error) {
	headNode := l.head
	if headNode == nil {
		return nil, errors.New("list is empty")
	}
	value := headNode.Value

	l.head = headNode.next
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}

	return value, nil
}

func (l *List) Pop() (interface{}, error) {
	tailNode := l.tail
	if tailNode == nil {
		return nil, errors.New("list is empty")
	}
	value := tailNode.Value

	l.tail = tailNode.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}

	return value, nil
}

func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = curr

	for curr != nil {
		next := curr.next
		curr.next = prev
		curr.prev = next

		prev = curr
		curr = next
	}

	l.head = prev
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
