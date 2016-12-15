package datastrs

import (
	"errors"
	"fmt"
	"strings"
)

var ErrLNotFound = errors.New("Element not found in Linked List")
var ErrLEmpty = errors.New("Linked List is empty")

// Structure of node
type Node struct {
	element interface{}
	next    *Node
}

// Structure of Linked List
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func NewNode(value interface{}) *Node {
	return &Node{value, nil}
}

func (list *LinkedList) Add(value interface{}) {
	newNode := NewNode(value)
	if list.tail == nil && list.head == nil {
		// If tail and head are nil then we are adding our first element
		list.tail = newNode
		list.head = newNode
		list.size = list.size + 1
	} else {
		// Join the new node in the last
		list.tail.next = newNode
		list.tail = newNode
		list.size = list.size + 1
	}
}

func (list *LinkedList) Search(value interface{}) (int, error) {
	if list.size == 0 {
		return 0, ErrLEmpty
	}
	var it *Node = list.head
	for index := 0; index < list.size; index++ {
		if it.element == value {
			return index, nil
		}
		it = it.next
	}
	return 0, ErrLNotFound
}

// Removes first instance of the input value
func (list *LinkedList) Remove(value interface{}) (bool, error) {
	if list.size == 0 {
		return false, ErrLEmpty
	}
	var current *Node = list.head
	var prev *Node = list.head
	for index := 0; index < list.size; index++ {
		if current.element == value {
			if index == 0 {
				list.head = current.next
				current = nil
				list.size = list.size - 1
				return true, nil
			} else if index == list.size-1 {
				list.tail = prev
				current = nil
				list.size = list.size - 1
				return true, nil
			} else {
				prev.next = current.next
				list.size = list.size - 1
				return true, nil
			}
		}
		prev = current
		current = current.next
	}
	return false, ErrLNotFound
}

func (list *LinkedList) String() string {
	var s string = "Linked List :"
	var values []string
	var it *Node = list.head
	for index := 0; index < list.size; index++ {
		values = append(values, fmt.Sprintf("%v", it.element))
		it = it.next
	}
	s += " [ " + strings.Join(values, ", ") + " ] "
	s += "having size"
	s += fmt.Sprintf(" %v\n", list.size)
	return s
}
