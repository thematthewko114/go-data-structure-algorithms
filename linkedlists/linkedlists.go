package linkedlists

import (
	"fmt"
)

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
	}
}

func (l *LinkedList) Add(value string) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (l *LinkedList) Remove(value string) {
	if l.head == nil {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		return
	}

	curr := l.head
	for curr.next != nil && curr.next.data != value {
		curr = curr.next
	}

	if curr.next != nil {
		curr.next = curr.next.next
	}
}

func (l *LinkedList) PrintList() {
	fmt.Print("Front |")
	curr := l.head
	for curr != nil {
		fmt.Printf("%s ", curr.data)
		curr = curr.next
	}
	fmt.Print("| Back")
	fmt.Println()
}
