package datastrs

import (
	"fmt"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	testingList := NewLinkedList()
	fmt.Println(testingList.String())

	// Output:
	// Linked List : [  ] having size 0
}

func TestLinkedList_Add(t *testing.T) {
	testingList := NewLinkedList()
	testingList.Add(1)
	testingList.Add(2)
	fmt.Println(testingList.String())

	// Output:
	// Linked List : [ 1, 2 ] having size 2
}

func TestLinkedList_Search(t *testing.T) {
	testingList := NewLinkedList()
	testingList.Add(1)
	testingList.Add(2)
	_, error := testingList.Search(3)
	if error != nil {
		fmt.Println(error)
	}
	value, error := testingList.Search(2)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Print("Search index -> " + fmt.Sprintf("%v\n", value))
	}
	fmt.Println(testingList.String())

	// Output:
	// Element not found in Linked List
	// Search index -> 1
	// Linked List : [ 1, 2 ] having size 2
}

func TestLinkedList_Remove(t *testing.T) {
	testingList := NewLinkedList()
	_, error := testingList.Remove("b")
	if error != nil {
		fmt.Println(error)
	}
	testingList.Add("a")
	testingList.Add("b")
	_, error = testingList.Remove(2)
	if error != nil {
		fmt.Println(error)
	}
	value, error := testingList.Remove("a")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Print("Node removed -> " + fmt.Sprintf("%v\n", value))
	}
	value, error = testingList.Remove("b")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Print("Node removed -> " + fmt.Sprintf("%v\n", value))
	}
	fmt.Println(testingList.String())

	// Output:
	// Linked List is empty
	// Element not found in Linked List
	// Node removed -> true
	// Node removed -> true
	// Linked List : [  ] having size 0
}
