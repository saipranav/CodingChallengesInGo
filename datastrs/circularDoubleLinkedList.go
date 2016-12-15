package datastrs

import (
	"errors"
	"fmt"
	"strings"
)

var ErrCDLNotFound = errors.New("Element not found in Cir Double Linked List")
var ErrCDLEmpty = errors.New("Cir Double Linked List is empty")

// Structure of node
type CDNode struct {
	element interface{}
	next    *CDNode
	prev    *CDNode
}

// Structure of Linked List
type CDLinkedList struct {
	head *CDNode
	tail *CDNode
	size int
}

func NewCirDoubleLinkedList() *CDLinkedList {
	return &CDLinkedList{}
}

func NewCDNode(value interface{}) *CDNode {
	return &CDNode{value, nil, nil}
}

func (list *CDLinkedList) Add(value interface{}) {
	newCDNode := NewCDNode(value)
	if list.tail == nil && list.head == nil {
		// If tail and head are nil then we are adding our first element
		list.tail = newCDNode
		list.head = newCDNode
		newCDNode.next = newCDNode
		newCDNode.prev = newCDNode
		list.size = list.size + 1
	} else {
		// Join the new node in the last
		newCDNode.prev = list.tail
		newCDNode.next = list.head
		list.tail.next = newCDNode
		list.tail = newCDNode
		list.size = list.size + 1
	}
}

func (list *CDLinkedList) Search(value interface{}) (int, error) {
	if list.size == 0 {
		return 0, ErrCDLEmpty
	}
	var it *CDNode = list.head
	for index := 0; index < list.size; index++ {
		if it.element == value {
			return index, nil
		}
		it = it.next
	}
	return 0, ErrCDLNotFound
}

// Removes first instance of the input value
func (list *CDLinkedList) Remove(value interface{}) (bool, error) {
	if list.size == 0 {
		return false, ErrCDLEmpty
	}
	var current *CDNode = list.head
	var prev *CDNode = list.head
	for index := 0; index < list.size; index++ {
		if current.element == value {
			if index == 0 {
				list.head = current.next
				list.head.prev = list.tail
				current = nil
				list.size = list.size - 1
				return true, nil
			} else if index == list.size-1 {
				list.tail = prev
				list.tail.next = list.head
				current = nil
				list.size = list.size - 1
				return true, nil
			} else {
				prev.next = current.next
				current.next.prev = prev
				list.size = list.size - 1
				return true, nil
			}
		}
		prev = current
		current = current.next
	}
	return false, ErrCDLNotFound
}

func (list *CDLinkedList) String() string {
	var s string = "Cir Double Linked List :"
	var values []string
	var it *CDNode = list.head
	for index := 0; index < list.size; index++ {
		values = append(values, fmt.Sprintf("%v", it.element))
		it = it.next
	}
	s += " [ " + strings.Join(values, ", ") + " ] "
	s += "having size"
	s += fmt.Sprintf(" %v\n", list.size)
	return s
}
